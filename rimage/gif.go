// Copyright (c) 2024 Barat Semet (https://github.com/barats)
// Resizem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package rimage

import (
	"bytes"
	"errors"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"io"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

// AnimatedGIF is implemented by images that can write themselves out as a
// multi-frame GIF. GIF sources keep all their frames; every other source
// produces a single-frame GIF.
type AnimatedGIF interface {
	WriteAnimatedGIFFile(dest string, colors int) error
}

type GIFImage struct {
	image  image.Image
	source string

	width  int
	height int

	// animation holds the source's frames when it is an animated GIF;
	// nil for single-frame GIFs.
	animation *gifAnimation
}

// gifAnimation carries the frames of an animated GIF through resize, so the
// animation survives conversion instead of being flattened to one frame.
type gifAnimation struct {
	frames []image.Image // display frames, one per source frame
	delays []int         // per-frame delay in 100ths of a second
	loops  int           // loop count, same semantics as image/gif.GIF.LoopCount
	width  int           // logical screen width before resize
	height int           // logical screen height before resize
}

func NewGIFImage(file io.Reader, w, h int, s string, autoOrientation bool) (*GIFImage, error) {
	anim, err := gif.DecodeAll(file)
	if err == nil && len(anim.Image) > 0 {
		img := &GIFImage{width: w, height: h, source: s, image: anim.Image[0]}
		if len(anim.Image) > 1 {
			img.animation = &gifAnimation{
				delays: append([]int(nil), anim.Delay...),
				loops:  anim.LoopCount,
				width:  anim.Config.Width,
				height: anim.Config.Height,
			}
			for _, frame := range anim.Image {
				img.animation.frames = append(img.animation.frames, frame)
			}
		}
		return img, nil
	}

	// Fall back to a single-frame decode for GIFs that gif.DecodeAll
	// rejects. The reader is rewound first when it supports seeking.
	if seeker, ok := file.(io.Seeker); ok {
		if _, seekErr := seeker.Seek(0, io.SeekStart); seekErr == nil {
			decoded, decodeErr := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
			if decodeErr == nil {
				return &GIFImage{width: w, height: h, source: s, image: decoded}, nil
			}
		}
	}
	return nil, err
}

func (i *GIFImage) Size() (int, int) {
	return i.width, i.height
}

func (i *GIFImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *GIFImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}

func (i *GIFImage) Type() OutputImageType {
	return GIF
}

func (i *GIFImage) Data() image.Image {
	return i.image
}

func (i *GIFImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)

	if i.animation != nil {
		anim, err := resizeAnimatedGIF(i.animation, w, h, filter)
		if err != nil {
			// Fall back to a static output rather than writing frames
			// that do not match the requested size.
			i.animation = nil
			return
		}
		i.animation = anim
	}
}

// WriteAnimatedGIFFile writes the image as a GIF. Animated sources keep all
// their frames (with original delays); everything else writes a single frame.
func (i *GIFImage) WriteAnimatedGIFFile(dest string, colors int) error {
	if i.animation == nil {
		return CreateGIFFile(dest, i.image, colors)
	}

	out := &gif.GIF{
		Delay:     i.animation.delays,
		Config:    image.Config{Width: i.animation.width, Height: i.animation.height},
		LoopCount: i.animation.loops,
	}
	for _, frame := range i.animation.frames {
		pf, err := frameToPaletted(frame, colors)
		if err != nil {
			return err
		}
		out.Image = append(out.Image, pf)
		// Every output frame is a full-size display frame, so the next
		// frame must not rely on disposing the previous one.
		out.Disposal = append(out.Disposal, gif.DisposalNone)
	}

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()
	return gif.EncodeAll(f, out)
}

// resizeAnimatedGIF scales every frame of an animated GIF. Each frame is
// resized as-is and placed at its proportionally scaled position on a
// full-size transparent canvas, so the output frames can be overlaid on each
// other by the viewer (disposal is set to none when writing).
func resizeAnimatedGIF(anim *gifAnimation, w, h int, filter imaging.ResampleFilter) (*gifAnimation, error) {
	if anim.width <= 0 || anim.height <= 0 {
		return nil, errors.New("gif: cannot resize GIF with unknown logical screen size")
	}
	if w <= 0 {
		w = anim.width
	}
	if h <= 0 {
		h = anim.height
	}
	sx := float64(w) / float64(anim.width)
	sy := float64(h) / float64(anim.height)

	out := &gifAnimation{
		delays: make([]int, len(anim.frames)),
		loops:  anim.loops,
		width:  w,
		height: h,
	}
	for i, frame := range anim.frames {
		out.delays[i] = anim.delays[i]

		b := frame.Bounds()
		nx := int(math.Round(float64(b.Min.X) * sx))
		ny := int(math.Round(float64(b.Min.Y) * sy))
		nw := int(math.Round(float64(b.Dx()) * sx))
		nh := int(math.Round(float64(b.Dy()) * sy))
		if nw < 1 {
			nw = 1
		}
		if nh < 1 {
			nh = 1
		}

		resized := imaging.Resize(frame, nw, nh, filter)
		canvas := image.NewRGBA(image.Rect(0, 0, w, h))
		draw.Draw(canvas, image.Rect(nx, ny, nx+nw, ny+nh), resized, resized.Bounds().Min, draw.Over)
		out.frames = append(out.frames, canvas)
	}
	return out, nil
}

// nearestDrawer maps each source pixel to the nearest palette color without
// error diffusion, so transparency never bleeds into neighboring pixels.
type nearestDrawer struct{}

func (nearestDrawer) Draw(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
	draw.Draw(dst, r, src, sp, draw.Src)
}

// transparentPlan9Quantizer builds a palette whose first entry is transparent,
// followed by the Plan9 palette. gif.Encode's default palette has no
// transparent entry, which would turn transparent frame areas opaque.
type transparentPlan9Quantizer struct {
	colors int
}

func (q transparentPlan9Quantizer) Quantize(p color.Palette, m image.Image) color.Palette {
	n := q.colors - 1
	if n > len(palette.Plan9) {
		n = len(palette.Plan9)
	}
	out := make(color.Palette, 0, n+1)
	out = append(out, color.RGBA{}) // transparent
	return append(out, palette.Plan9[:n]...)
}

// frameToPaletted quantizes a full-size frame to a paletted image for GIF
// encoding. Transparency is preserved via the palette's transparent entry.
func frameToPaletted(img image.Image, colors int) (*image.Paletted, error) {
	if colors <= 0 || colors > 256 {
		colors = 256
	}
	var buf bytes.Buffer
	opts := &gif.Options{
		NumColors: colors,
		Drawer:    nearestDrawer{},
		Quantizer: transparentPlan9Quantizer{colors: colors},
	}
	if err := gif.Encode(&buf, img, opts); err != nil {
		return nil, err
	}
	decoded, err := gif.Decode(&buf)
	if err != nil {
		return nil, err
	}
	pf, ok := decoded.(*image.Paletted)
	if !ok {
		return nil, errors.New("gif: failed to quantize frame to paletted image")
	}
	return pf, nil
}

func CreateGIFFile(dest string, data image.Image, colors int) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	return imaging.Encode(out, data, imaging.GIF, imaging.GIFNumColors(colors))
}

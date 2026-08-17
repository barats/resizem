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
	"image"
	"image/color"
	"image/gif"
	"os"
	"path/filepath"
	"testing"
)

var gifTestPalette = color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xff}, // index 0: black
	color.RGBA{0xff, 0x00, 0x00, 0xff}, // index 1: red
	color.RGBA{0x00, 0xff, 0x00, 0xff}, // index 2: green
}

// writeAnimatedGIF writes a two-frame GIF to path. Frame 0 is red, frame 1 is
// green, both full-screen at w x h.
func writeAnimatedGIF(t *testing.T, path string, w, h int) {
	t.Helper()
	f0 := image.NewPaletted(image.Rect(0, 0, w, h), gifTestPalette)
	for i := range f0.Pix {
		f0.Pix[i] = 1
	}
	f1 := image.NewPaletted(image.Rect(0, 0, w, h), gifTestPalette)
	for i := range f1.Pix {
		f1.Pix[i] = 2
	}
	anim := &gif.GIF{
		Image:     []*image.Paletted{f0, f1},
		Delay:     []int{10, 20},
		LoopCount: 0, // loop forever
		Disposal:  []byte{gif.DisposalNone, gif.DisposalNone},
		Config:    image.Config{ColorModel: gifTestPalette, Width: w, Height: h},
	}
	f, err := os.Create(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := gif.EncodeAll(f, anim); err != nil {
		t.Fatal(err)
	}
}

func decodeGIF(t *testing.T, path string) *gif.GIF {
	t.Helper()
	f, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	g, err := gif.DecodeAll(f)
	if err != nil {
		t.Fatalf("DecodeAll(%s): %v", path, err)
	}
	return g
}

func TestAnimatedGIFResizePreservesFrames(t *testing.T) {
	src := filepath.Join(t.TempDir(), "anim.gif")
	writeAnimatedGIF(t, src, 20, 10)

	img, err := CreateImage(src, false)
	if err != nil {
		t.Fatalf("CreateImage: %v", err)
	}
	gi, ok := img.(*GIFImage)
	if !ok {
		t.Fatalf("CreateImage returned %T, want *GIFImage", img)
	}
	if gi.animation == nil {
		t.Fatal("expected animated GIF frames to be kept")
	}

	img.Resize(40, 20, Lanczos)

	dest := filepath.Join(t.TempDir(), "out.gif")
	if err := gi.WriteAnimatedGIFFile(dest, 256); err != nil {
		t.Fatalf("WriteAnimatedGIFFile: %v", err)
	}

	out := decodeGIF(t, dest)
	if len(out.Image) != 2 {
		t.Fatalf("output frame count = %d, want 2", len(out.Image))
	}
	if out.Config.Width != 40 || out.Config.Height != 20 {
		t.Errorf("output size = %dx%d, want 40x20", out.Config.Width, out.Config.Height)
	}
	if out.Delay[0] != 10 || out.Delay[1] != 20 {
		t.Errorf("output delays = %v, want [10 20]", out.Delay)
	}
	if out.LoopCount != 0 {
		t.Errorf("output LoopCount = %d, want 0", out.LoopCount)
	}
	if len(out.Disposal) != 2 {
		t.Errorf("output disposal length = %d, want 2", len(out.Disposal))
	}

	center := image.Pt(20, 10)
	r0, g0, _, _ := out.Image[0].At(center.X, center.Y).RGBA()
	r1, g1, _, _ := out.Image[1].At(center.X, center.Y).RGBA()
	if r0 <= g0 {
		t.Errorf("frame 0 center not red-ish (r=%d g=%d)", r0>>8, g0>>8)
	}
	if g1 <= r1 {
		t.Errorf("frame 1 center not green-ish (r=%d g=%d)", r1>>8, g1>>8)
	}
}

func TestAnimatedGIFDeltaFramePositioning(t *testing.T) {
	// Frame 0 fills the screen; frame 1 is a green sub-region that must be
	// scaled and placed at the proportional position on the output canvas.
	f0 := image.NewPaletted(image.Rect(0, 0, 20, 10), gifTestPalette)
	for i := range f0.Pix {
		f0.Pix[i] = 1
	}
	f1 := image.NewPaletted(image.Rect(5, 2, 15, 7), gifTestPalette)
	for i := range f1.Pix {
		f1.Pix[i] = 2
	}
	anim := &gif.GIF{
		Image:     []*image.Paletted{f0, f1},
		Delay:     []int{10, 10},
		LoopCount: 0,
		Disposal:  []byte{gif.DisposalNone, gif.DisposalNone},
		Config:    image.Config{ColorModel: gifTestPalette, Width: 20, Height: 10},
	}
	src := filepath.Join(t.TempDir(), "delta.gif")
	f, err := os.Create(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := gif.EncodeAll(f, anim); err != nil {
		t.Fatal(err)
	}
	f.Close()

	img, err := CreateImage(src, false)
	if err != nil {
		t.Fatalf("CreateImage: %v", err)
	}
	gi := img.(*GIFImage)
	img.Resize(40, 20, NearestNeighbor)

	dest := filepath.Join(t.TempDir(), "out.gif")
	if err := gi.WriteAnimatedGIFFile(dest, 256); err != nil {
		t.Fatalf("WriteAnimatedGIFFile: %v", err)
	}

	out := decodeGIF(t, dest)
	if len(out.Image) != 2 {
		t.Fatalf("output frame count = %d, want 2", len(out.Image))
	}
	// Green delta scaled by 2 lands at (10,4)-(30,14).
	_, g, _, a := out.Image[1].At(20, 9).RGBA()
	if a == 0 || g == 0 {
		t.Errorf("delta interior not green and opaque (a=%d g=%d)", a>>8, g>>8)
	}
	_, _, _, a = out.Image[1].At(0, 0).RGBA()
	if a != 0 {
		t.Errorf("delta exterior should be transparent, got alpha %d", a>>8)
	}
	if out.Disposal[1] != gif.DisposalNone {
		t.Errorf("frame 1 disposal = %d, want DisposalNone", out.Disposal[1])
	}
}

func TestSingleFrameGIFWritesStaticOutput(t *testing.T) {
	m := image.NewPaletted(image.Rect(0, 0, 10, 10), gifTestPalette)
	for i := range m.Pix {
		m.Pix[i] = 1
	}
	anim := &gif.GIF{
		Image:  []*image.Paletted{m},
		Delay:  []int{0},
		Config: image.Config{ColorModel: gifTestPalette, Width: 10, Height: 10},
	}
	src := filepath.Join(t.TempDir(), "single.gif")
	f, err := os.Create(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := gif.EncodeAll(f, anim); err != nil {
		t.Fatal(err)
	}
	f.Close()

	img, err := CreateImage(src, false)
	if err != nil {
		t.Fatalf("CreateImage: %v", err)
	}
	gi := img.(*GIFImage)
	if gi.animation != nil {
		t.Fatal("single-frame GIF must not be treated as animated")
	}

	img.Resize(20, 20, Lanczos)
	dest := filepath.Join(t.TempDir(), "out.gif")
	if err := gi.WriteAnimatedGIFFile(dest, 256); err != nil {
		t.Fatalf("WriteAnimatedGIFFile: %v", err)
	}

	out := decodeGIF(t, dest)
	if len(out.Image) != 1 {
		t.Errorf("output frame count = %d, want 1", len(out.Image))
	}
	if out.Config.Width != 20 || out.Config.Height != 20 {
		t.Errorf("output size = %dx%d, want 20x20", out.Config.Width, out.Config.Height)
	}
}

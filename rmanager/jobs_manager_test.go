// Copyright (c) 2024 Barat Semet (https://github.com/barats)
// Resizem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package rmanager

import (
	"image"
	"image/color"
	"image/gif"
	"os"
	"path/filepath"
	"testing"

	"resizem/rimage"
)

func writeTwoFrameGIF(t *testing.T, path string) {
	t.Helper()
	pal := color.Palette{
		color.RGBA{0x00, 0x00, 0x00, 0xff},
		color.RGBA{0xff, 0x00, 0x00, 0xff},
		color.RGBA{0x00, 0xff, 0x00, 0xff},
	}
	f0 := image.NewPaletted(image.Rect(0, 0, 20, 10), pal)
	for i := range f0.Pix {
		f0.Pix[i] = 1
	}
	f1 := image.NewPaletted(image.Rect(0, 0, 20, 10), pal)
	for i := range f1.Pix {
		f1.Pix[i] = 2
	}
	anim := &gif.GIF{
		Image:     []*image.Paletted{f0, f1},
		Delay:     []int{10, 20},
		LoopCount: 0,
		Disposal:  []byte{gif.DisposalNone, gif.DisposalNone},
		Config:    image.Config{ColorModel: pal, Width: 20, Height: 10},
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

// TestWriteFileByFormatPreservesAnimatedGIF exercises the real write path a
// job would use: converting an animated GIF to a resized GIF keeps every frame.
func TestWriteFileByFormatPreservesAnimatedGIF(t *testing.T) {
	src := filepath.Join(t.TempDir(), "anim.gif")
	writeTwoFrameGIF(t, src)

	img, err := rimage.CreateImage(src, false)
	if err != nil {
		t.Fatalf("CreateImage: %v", err)
	}
	img.Resize(40, 20, rimage.Lanczos)

	dest, err := writeFileByFormat(t.TempDir(), img, 40, 20,
		rimage.ImageOptions{Format: rimage.GIF, GIFNumColors: 256})
	if err != nil {
		t.Fatalf("writeFileByFormat: %v", err)
	}

	f, err := os.Open(dest)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	out, err := gif.DecodeAll(f)
	if err != nil {
		t.Fatalf("DecodeAll output: %v", err)
	}
	if len(out.Image) != 2 {
		t.Errorf("output frame count = %d, want 2", len(out.Image))
	}
	if out.Config.Width != 40 || out.Config.Height != 20 {
		t.Errorf("output size = %dx%d, want 40x20", out.Config.Width, out.Config.Height)
	}
	if out.Delay[0] != 10 || out.Delay[1] != 20 {
		t.Errorf("output delays = %v, want [10 20]", out.Delay)
	}
}

// TestWriteFileByFormatNonGIFSourceToGIF keeps the pre-existing behavior of
// converting a still image into a single-frame GIF.
func TestWriteFileByFormatNonGIFSourceToGIF(t *testing.T) {
	src := filepath.Join(t.TempDir(), "still.png")
	png := image.NewRGBA(image.Rect(0, 0, 10, 10))
	f, err := os.Create(src)
	if err != nil {
		t.Fatal(err)
	}
	if err := rimage.CreatePNGFile(src, png, -1); err != nil {
		t.Fatal(err)
	}
	f.Close()

	img, err := rimage.CreateImage(src, false)
	if err != nil {
		t.Fatalf("CreateImage: %v", err)
	}
	img.Resize(20, 20, rimage.Lanczos)

	dest, err := writeFileByFormat(t.TempDir(), img, 20, 20,
		rimage.ImageOptions{Format: rimage.GIF, GIFNumColors: 256})
	if err != nil {
		t.Fatalf("writeFileByFormat: %v", err)
	}

	rf, err := os.Open(dest)
	if err != nil {
		t.Fatal(err)
	}
	defer rf.Close()
	out, err := gif.DecodeAll(rf)
	if err != nil {
		t.Fatalf("DecodeAll output: %v", err)
	}
	if len(out.Image) != 1 {
		t.Errorf("output frame count = %d, want 1", len(out.Image))
	}
	if out.Config.Width != 20 || out.Config.Height != 20 {
		t.Errorf("output size = %dx%d, want 20x20", out.Config.Width, out.Config.Height)
	}
}

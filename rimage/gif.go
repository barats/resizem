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
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
)

type GIFImage struct {
	image  image.Image
	source string

	width  int
	height int
}

func NewGIFImage(file io.Reader, w, h int, s string, autoOrientation bool) (*GIFImage, error) {
	//img, err := gif.Decode(file)
	img, err := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
	if err != nil {
		return nil, err
	}

	return &GIFImage{width: w, height: h, source: s, image: img}, nil
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
}

func CreateGIFFile(dest string, data image.Image, colors int) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	//return gif.Encode(out, data, nil)
	return imaging.Encode(out, data, imaging.GIF, imaging.GIFNumColors(colors))
}

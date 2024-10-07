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

type BMPImage struct {
	image  image.Image
	width  int
	height int
	source string
}

func NewBMPImage(file io.Reader, w, h int, s string, autoOrientation bool) (*BMPImage, error) {
	//img, err := bmp.Decode(file)
	img, err := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
	imaging.Decode(file)
	if err != nil {
		return nil, err
	}

	return &BMPImage{width: w, height: h, source: s, image: img}, nil
}

func (i *BMPImage) Type() OutputImageType {
	return BMP
}

func (i *BMPImage) Size() (int, int) {
	return i.width, i.height
}

func (i *BMPImage) Data() image.Image {
	return i.image
}

func (i *BMPImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *BMPImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}

func (i *BMPImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}

func CreateBMPFile(dest string, data image.Image) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	//return bmp.Encode(out, data)
	return imaging.Encode(out, data, imaging.BMP)
}

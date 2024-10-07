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
	"golang.org/x/image/tiff"
)

type TIFFImage struct {
	image  image.Image
	source string
	width  int
	height int
}

func NewTIFFImage(file io.Reader, w, h int, s string, autoOrientation bool) (*TIFFImage, error) {
	//img, err := tiff.Decode(file)
	img, err := imaging.Decode(file, imaging.AutoOrientation(autoOrientation))
	if err != nil {
		return nil, err
	}

	return &TIFFImage{width: w, height: h, source: s, image: img}, nil
}

func (i *TIFFImage) Size() (int, int) {
	return i.width, i.height
}

func (i *TIFFImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *TIFFImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}

func (i *TIFFImage) Data() image.Image {
	return i.image
}

func (i *TIFFImage) Type() OutputImageType {
	return TIFF
}

func (i *TIFFImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}

func CreateTIFFFile(dest string, data image.Image, compression int) error {
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	return tiff.Encode(out, data, &tiff.Options{Compression: tiff.CompressionType(compression)})
}

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
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"golang.org/x/image/webp"
)

type WebPImage struct {
	image  image.Image
	source string
	width  int
	height int
}

func NewWebPImage(file io.Reader, w, h int, s string) (*WebPImage, error) {
	img, err := webp.Decode(file)
	if err != nil {
		return nil, err
	}

	return &WebPImage{width: w, height: h, source: s, image: img}, nil
}

func (i *WebPImage) Size() (int, int) {
	return i.width, i.height
}

func (i *WebPImage) Type() OutputImageType {
	return WEBP
}

func (i *WebPImage) Dir() string {
	return filepath.Dir(i.source)
}

func (i *WebPImage) Name() string {
	return strings.TrimSuffix(filepath.Base(i.source), filepath.Ext(i.source))
}
func (i *WebPImage) Data() image.Image {
	return i.image
}

func (i *WebPImage) Resize(w, h int, f ResampleFilterType) {
	filter := MatchFilter(f)

	if w <= 0 {
		w = i.width
	}
	if h <= 0 {
		h = i.height
	}

	i.image = imaging.Resize(i.image, w, h, filter)
}

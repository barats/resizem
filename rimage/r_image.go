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
	"fmt"
	"image"
	"os"
)

type ResizemImage interface {
	//Size() return width, height in order
	Size() (int, int)

	//Data() returns decoded image data
	Data() image.Image

	//Dir() returns all but the last element of path, TYPICALLY the path's directory.
	Dir() string

	// Name() returns file name WITHOUT extension
	Name() string

	//Resize image to fit given width and height with filter
	Resize(width, height int, filter ResampleFilterType)

	//Type() returns image type
	Type() OutputImageType
}

// CreateImage
//
// Create ResizemImage object from given source(absolute path & name of file)
//
// It also reads given file's properties and retrieve image format, width, height
func CreateImage(source string, autoOrientation bool) (ResizemImage, error) {
	file, err := os.Open(source)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	//Decode config, tries to retrieve image format & other properties
	config, format, err := image.DecodeConfig(file)
	if err != nil {
		return nil, err
	}

	//Seek back to (0,0) position of image file
	//Reason for this:  https://github.com/golang/go/issues/50992
	file.Seek(0, 0)

	//Try to create correct image
	switch format {
	case "png":
		return NewPNGImage(file, config.Width, config.Height, source, autoOrientation)
	case "jpg":
		return NewJPGImage(file, config.Width, config.Height, source, autoOrientation)
	case "jpeg":
		return NewJPGImage(file, config.Width, config.Height, source, autoOrientation)
	case "tif":
		return NewTIFFImage(file, config.Width, config.Height, source, autoOrientation)
	case "tiff":
		return NewTIFFImage(file, config.Width, config.Height, source, autoOrientation)
	case "webp":
		return NewWebPImage(file, config.Width, config.Height, source)
	case "bmp":
		return NewBMPImage(file, config.Width, config.Height, source, autoOrientation)
	case "gif":
		return NewGIFImage(file, config.Width, config.Height, source, autoOrientation)
	}

	return nil, fmt.Errorf("unsupported image format : %s", format)
}

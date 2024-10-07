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

import "github.com/disintegration/imaging"

type OutputImageType int

const (
	JPG OutputImageType = iota + 1 // 1
	JPEG
	PNG
	BMP
	GIF
	TIFF
	WEBP
)

type SupportedOutputImageType struct {
	Name  string          `json:"name"`
	Value OutputImageType `json:"value"`
}

var OutputImagesTypes = []SupportedOutputImageType{
	{Name: "JPG", Value: JPG},
	{Name: "JPEG", Value: JPEG},
	{Name: "PNG", Value: PNG},
	{Name: "BMP", Value: BMP},
	{Name: "GIF", Value: GIF},
	{Name: "TIFF", Value: TIFF},
}

type ImageOptions struct {
	Format          OutputImageType    `json:"dest_format"`
	Filter          ResampleFilterType `json:"resample_filter"`
	Path            string             `json:"desc_path"`
	Width           int                `json:"dest_width"`
	Height          int                `json:"dest_height"`
	JPEGQuality     int                `json:"jpeg_quality"`
	GIFNumColors    int                `json:"gif_number_of_colors"`
	TIFFCompression int                `json:"tiff_compression"`
	PNGCompression  int                `json:"png_compression"`
	AutoOrientation bool               `json:"auto_orientation"`
	CPUMemUsage     int                `json:"cpu_memory_usage"`
}

// Website: https://github.com/disintegration/imaging

//Imaging supports image resizing using various resampling filters. The most notable ones:
// Lanczos - A high-quality resampling filter for photographic images yielding sharp results.
// CatmullRom - A sharp cubic filter that is faster than Lanczos filter while providing similar results.
// MitchellNetravali - A cubic filter that produces smoother results with less ringing artifacts than CatmullRom.
// Linear - Bilinear resampling filter, produces smooth output. Faster than cubic filters.
// Box - Simple and fast averaging filter appropriate for downscaling. When upscaling it's similar to NearestNeighbor.
// NearestNeighbor - Fastest resampling filter, no antialiasing.

//NearestNeighbor, Box, Linear, Hermite, MitchellNetravali, CatmullRom, BSpline, Gaussian, Lanczos, Hann, Hamming, Blackman, Bartlett, Welch, Cosine.

type ResampleFilterType int

const (
	Lanczos ResampleFilterType = iota + 1 //1
	CatmullRom
	MitchellNetravali
	Linear
	Box
	NearestNeighbor
	Hermite
	BSpline
	Gaussian
	Hann
	Hamming
	Blackman
	Bartlett
	Welch
	Cosine
)

type SupportedResampleFilterType struct {
	Name  string             `json:"name"`
	Value ResampleFilterType `json:"value"`
}

var ResampleFilterTypes = []SupportedResampleFilterType{
	{Name: "Lanczos", Value: Lanczos},
	{Name: "CatmullRom", Value: CatmullRom},
	{Name: "MitchellNetravali", Value: MitchellNetravali},
	{Name: "Linear", Value: Linear},
	{Name: "Box", Value: Box},
	{Name: "NearestNeighbor", Value: NearestNeighbor},
	{Name: "Hermite", Value: Hermite},
	{Name: "BSpline", Value: BSpline},
	{Name: "Gaussian", Value: Gaussian},
	{Name: "Hann", Value: Hann},
	{Name: "Hamming", Value: Hamming},
	{Name: "Blackman", Value: Blackman},
	{Name: "Bartlett", Value: Bartlett},
	{Name: "Welch", Value: Welch},
	{Name: "Cosine", Value: Cosine},
}

func MatchFilter(value ResampleFilterType) imaging.ResampleFilter {
	switch value {
	case Lanczos:
		return imaging.Lanczos
	case CatmullRom:
		return imaging.CatmullRom
	case MitchellNetravali:
		return imaging.MitchellNetravali
	case Linear:
		return imaging.Linear
	case Box:
		return imaging.Box
	case NearestNeighbor:
		return imaging.NearestNeighbor
	case Hermite:
		return imaging.Hermite
	case BSpline:
		return imaging.BSpline
	case Gaussian:
		return imaging.Gaussian
	case Hann:
		return imaging.Hann
	case Hamming:
		return imaging.Hamming
	case Blackman:
		return imaging.Blackman
	case Bartlett:
		return imaging.Bartlett
	case Welch:
		return imaging.Welch
	case Cosine:
		return imaging.Cosine
	}
	return imaging.Lanczos
}

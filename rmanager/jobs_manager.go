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
	"context"
	"fmt"
	"image"
	"log"
	"math"
	"path/filepath"
	"resizem/rimage"
	"runtime"
	"strings"

	"github.com/remeh/sizedwaitgroup"
)

type FileResult struct {
	Name    string `json:"name"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var jobContext, cancel = context.WithCancel(context.Background())

func CancelJobs() {
	cancel()
	jobContext.Done()
}

func StartJobs(ctx context.Context, files []string, opts rimage.ImageOptions) {
	cpu := runtime.NumCPU()
	workerCount := math.Round(float64(cpu) * float64(opts.CPUMemUsage))
	swg := sizedwaitgroup.New(int(workerCount)) //Should be better, if we could think about memory limit
	for _, f := range files {
		if err := swg.AddWithContext(jobContext); err != nil {
			log.Println(err)
			cancel()
		}
		go dealWithFile(ctx, &swg, f, opts)
	}
	swg.Wait()
}

func dealWithFile(ctx context.Context, wg *sizedwaitgroup.SizedWaitGroup,
	file string, option rimage.ImageOptions) {
	defer wg.Done()

	logMsg := fmt.Sprintf("file: %s, options: %v", file, option)
	log.Println(logMsg)

	//Create image instance
	img, err := rimage.CreateImage(file, option.AutoOrientation)
	if err != nil {
		//Event: error
		log.Println(err)
		msg := FileResult{Name: file, Status: -1, Message: err.Error()}
		SendFileResultEvent(ctx, msg)
		return
	}

	// Get file path
	// If there's any given destination path, use it instead
	dPath := img.Dir()
	if !strings.EqualFold(option.Path, "") {
		dPath = option.Path
	}

	// Get file size
	// If there's given width and height
	// and ONLY IF given parameters are CORRECT, use them instead
	dw, dh := img.Size()
	if option.Width > 0 {
		dw = option.Width
	}
	if option.Height > 0 {
		dh = option.Height
	}

	// If there's NO given output format
	// Then set image's original format to output format
	if option.Format == 0 {
		option.Format = img.Type()
	}

	//Do resize job
	img.Resize(dw, dh, option.Filter)

	//Write file by format
	rs, err := writeFileByFormat(dPath, img.Name(), img.Data(), dw, dh, option)
	if err != nil {
		//Event: error
		log.Println(err)
		msg := FileResult{Name: rs, Status: -1, Message: err.Error()}
		SendFileResultEvent(ctx, msg)
		return
	}

	//Event: success
	msg := FileResult{Name: rs, Status: 1}
	SendFileResultEvent(ctx, msg)
}

func writeFileByFormat(path, name string, data image.Image,
	width, height int, option rimage.ImageOptions) (string, error) {
	file := fmt.Sprintf("%s_%dx%d", name, width, height)
	dest := filepath.Join(path, file)
	switch option.Format {
	case rimage.BMP:
		return dest + ".bmp", rimage.CreateBMPFile(dest+".bmp", data)
	case rimage.GIF:
		return dest + ".gif", rimage.CreateGIFFile(dest+".gif", data, option.GIFNumColors)
	case rimage.JPEG:
		return dest + ".jpeg", rimage.CreateJPEGFile(dest+".jpeg", data, option.JPEGQuality)
	case rimage.JPG:
		return dest + ".jpg", rimage.CreateJPEGFile(dest+".jpg", data, option.JPEGQuality)
	case rimage.PNG:
		return dest + ".png", rimage.CreatePNGFile(dest+".png", data, option.PNGCompression)
	case rimage.TIFF:
		return dest + ".tiff", rimage.CreateTIFFFile(dest+".tiff", data, option.TIFFCompression)
	case rimage.WEBP:
		//Webp will be transformed to PNG, since I could not find a proper encoder for WebP
		return dest + ".png", rimage.CreatePNGFile(dest+".png", data, option.PNGCompression)
	}
	return "", nil
}

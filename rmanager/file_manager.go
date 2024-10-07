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
	"log"
	"os"
	"path/filepath"
	"resizem/rimage"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// Listening for DnD events
func FileDropEventListener(ctx context.Context) {
	runtime.OnFileDrop(ctx, func(x, y int, paths []string) {
		for _, p := range paths {
			if hasImageExt(p) {
				SendFilePathEvent(ctx, p)
			}
		} //end of for
	})
}

///////////////////////// FileManager /////////////////////////

type FileManager struct {
	ctx context.Context
}

func NewFileManager() *FileManager {
	return &FileManager{}
}

func (fm *FileManager) OnStartup(ctx context.Context) {
	fm.ctx = ctx
}

// Accepts request from frontend to cancel running jobs
func (fm *FileManager) CancelHandleFiles() {
	CancelJobs()
}

// Accepts files from fronted and start to handle them
func (fm *FileManager) StartHandleFiles(files []string, opts rimage.ImageOptions) {
	if len(files) > 0 {
		StartJobs(fm.ctx, files, opts)
	}
}

func (fm *FileManager) OpenFilesDialog() {
	filter := runtime.FileFilter{
		DisplayName: "Image Files",
		Pattern:     "*.jpg;*.jpeg;*.png;*.webp;*.tif;*.tiff;*.bmp;*.gif",
	}

	home, _ := os.UserHomeDir()
	if strings.EqualFold(runtime.Environment(fm.ctx).Platform, "windows") {
		// Known issue
		// https://github.com/wailsapp/wails/issues/1381
		home = ""
	}

	files, _ := runtime.OpenMultipleFilesDialog(fm.ctx, runtime.OpenDialogOptions{
		Title:            Resizem.Name,
		DefaultDirectory: home,
		Filters:          []runtime.FileFilter{filter},
	})

	for _, p := range files {
		SendFilePathEvent(fm.ctx, p)
	} //end of for
}

func (fm *FileManager) OpenDirectoryDialog() string {
	home, _ := os.UserHomeDir()
	path, _ := runtime.OpenDirectoryDialog(fm.ctx, runtime.OpenDialogOptions{
		Title:            Resizem.Name,
		DefaultDirectory: home,
	})

	return path
}

// Simply, check image format by its extension
// We will do the REAL type-check when worker goroutine starts the job
func hasImageExt(path string) bool {
	ext := filepath.Ext(path)
	log.Println("ext --> " + ext)
	if len(ext) >= 3 { // file extension should be long enough, otherwise there's no point to check for it
		ext = strings.ToLower(ext)
		if strings.EqualFold(".bmp", ext) || strings.EqualFold(".gif", ext) ||
			strings.EqualFold(".jpg", ext) || strings.EqualFold(".jpeg", ext) ||
			strings.EqualFold(".png", ext) ||
			strings.EqualFold(".tif", ext) || strings.EqualFold(".tiff", ext) ||
			strings.EqualFold(".webp", ext) {
			return true
		}
		return false
	}
	return false
}

// Copyright (c) 2024 Barat Semet (https://github.com/barats)
// Resizem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

import (
	"context"
	"embed"
	"fmt"
	"resizem/rmanager"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var frontend embed.FS

//go:embed logo.png
var logo []byte

//go:embed wails.json
var wailsJson string

func main() {

	am := rmanager.NewAppManager(wailsJson, logo)
	fm := rmanager.NewFileManager()
	tm := rmanager.NewTypeManager()

	err := wails.Run(&options.App{
		Title:                            rmanager.Resizem.Name,
		Width:                            1100,
		Height:                           680,
		MinWidth:                         750,
		MinHeight:                        450,
		Frameless:                        false,
		LogLevel:                         logger.DEBUG,
		LogLevelProduction:               logger.DEBUG,
		ErrorFormatter:                   func(err error) any { return err.Error() },
		CSSDragProperty:                  "--wails-draggable",
		CSSDragValue:                     "drag",
		EnableDefaultContextMenu:         false,
		EnableFraudulentWebsiteDetection: false,
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
			//set DisableWebViewDrop to true will fail on Win10
			//https://github.com/wailsapp/wails/issues/3782
			DisableWebViewDrop: false,
			CSSDropProperty:    "--wails-drop-target",
			CSSDropValue:       "drop",
		},
		Windows: &windows.Options{
			Theme:                windows.Light,
			WindowIsTranslucent:  false,
			WebviewIsTransparent: false,
			WebviewUserDataPath:  "",
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            true,
			},
			Appearance:           mac.NSAppearanceNameAqua,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  true,
			About: &mac.AboutInfo{
				Title:   rmanager.Resizem.Name,
				Message: fmt.Sprintf("Version: %s\n%s \n\n © %s %d", rmanager.Resizem.Version, rmanager.Resizem.Comments, rmanager.Resizem.Copyright, time.Now().Year()),
				Icon:    rmanager.Resizem.Logo,
			},
		},
		Linux: &linux.Options{
			Icon:                rmanager.Resizem.Logo,
			WindowIsTranslucent: true,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         rmanager.Resizem.Name,
		},
		AssetServer: &assetserver.Options{
			Assets: frontend,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			am.OnStartup(ctx)
			fm.OnStartup(ctx)
			tm.OnStartup(ctx)
			rmanager.FileDropEventListener(ctx)
		},
		OnShutdown: func(ctx context.Context) {
			am.OnShutdown()
		},
		OnBeforeClose: func(ctx context.Context) (prevent bool) {
			return am.OnBeforeClose()
		},
		OnDomReady: func(ctx context.Context) {
			// rmanager.FileDropEventListener(ctx)
		},
		Bind: []interface{}{
			am, fm, tm,
		}, EnumBind: []interface{}{},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

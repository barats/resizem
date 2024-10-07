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

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func SendBeforeExitEvent(ctx context.Context, data interface{}) {
	emit(ctx, eventBeforeExit(), data)
}

func SendFilePathEvent(ctx context.Context, path string) {
	emit(ctx, eventFileDrop(), path)
}

func SendFileResultEvent(ctx context.Context, data interface{}) {
	emit(ctx, eventFileResult(), data)
}

// /////////////// private functions /////////////////
func emit(ctx context.Context, event string, data interface{}) {
	runtime.EventsEmit(ctx, event, data)
}

func eventBeforeExit() string {
	return fmt.Sprintf("resizem-%s-before-exit", Resizem.Version)
}

func eventFileDrop() string {
	return fmt.Sprintf("resizem-%s-file-drop", Resizem.Version)
}

func eventFileResult() string {
	return fmt.Sprintf("resizem-%s-file-result", Resizem.Version)
}

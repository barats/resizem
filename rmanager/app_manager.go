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

	"github.com/tidwall/gjson"
)

var (
	Resizem AppInfo
)

type AppInfo struct {
	Name      string
	Version   string
	Copyright string
	Comments  string
	Logo      []byte
}

type AppManager struct {
	ctx context.Context
}

func NewAppManager(wailsJson string, logo []byte) *AppManager {

	Resizem = AppInfo{}
	Resizem.Name = gjson.Get(wailsJson, "info.productName").String()
	Resizem.Version = gjson.Get(wailsJson, "info.productVersion").String()
	Resizem.Copyright = gjson.Get(wailsJson, "info.copyright").String()
	Resizem.Comments = gjson.Get(wailsJson, "info.comments").String()
	Resizem.Logo = logo

	return &AppManager{}
}

func (am *AppManager) OnStartup(ctx context.Context) {
	am.ctx = ctx
}

func (am *AppManager) OnBeforeClose() bool {
	//Tell frontend to do clean-up-jobs
	SendBeforeExitEvent(am.ctx, "do exit")
	return false
}

func (am *AppManager) OnShutdown() {
	//Frontend has been destroyed, need to cleanup some mess.
}

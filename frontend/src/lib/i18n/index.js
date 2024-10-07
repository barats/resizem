// Copyright (c) 2024 Barat Semet (https://github.com/barats)
// Resizem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details. 

import { browser } from '$app/environment';
import { init, register, locale } from 'svelte-i18n';

const locale_storage_key = 'user-locale';

function getLocalStore(key) {
	if (browser) return localStorage.getItem(key);
}

function setLocalStore(key, value) {
	if (browser) localStorage.setItem(key, value);
}

let userLocale = getLocalStore(locale_storage_key) || 'en-US';

export let appLocales = [
	{ value: 'en-US', name: 'English' },
	{ value: 'zh-CN', name: '简体中文' },
	{ value: 'zh-TW', name: '繁體中文' }
];

register('en-US', () => import('./locales/en_US.json'));
register('zh-CN', () => import('./locales/zh_CN.json'));
register('zh-TW', () => import('./locales/zh_TW.json'));

init({
	fallbackLocale: userLocale,
	initialLocale: userLocale
});

export function setAppLocale(value) {
	setLocalStore(locale_storage_key, value);
	locale.set(value);
}

export function getAppLocale() {
	return getLocalStore(locale_storage_key) || 'en-US';
}

// Copyright (c) 2024 Barat Semet (https://github.com/barats)
// Resizem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

import { writable } from '@macfja/svelte-persistent-store';
import {
	KEY_ASK_WHERE_TO_SAVE,
	KEY_CPU_VALUE,
	KEY_DOING,
	KEY_EXIF_ORIENTATION_VALUE,
	KEY_FILES_LIST,
	KEY_FILTER_VALUE,
	KEY_FORMAT_VALUE,
	KEY_GIF_COLORS_VALUE,
	KEY_HEIGHT_VALUE,
	KEY_JPEG_QUALITY_VALUE,
	KEY_PNG_COMPRESSION_VALUE,
	KEY_RESULT_LIST,
	KEY_TIFF_COMPRESSION_VALUE,
	KEY_WIDTH_VALUE
} from './app_consts';

export let doing = writable(KEY_DOING, false);
export let askWheretoSave = writable(KEY_ASK_WHERE_TO_SAVE, false);

export let filesList = writable(KEY_FILES_LIST, []);
export let resultList = writable(KEY_RESULT_LIST, []);

export let formatValue = writable(KEY_FORMAT_VALUE, 0); //Keep original format
export let filterValue = writable(KEY_FILTER_VALUE, 1); // Lanczos
export let widthValue = writable(KEY_WIDTH_VALUE);
export let heightValue = writable(KEY_HEIGHT_VALUE);

export let cpuUsageValue = writable(KEY_CPU_VALUE, 1);
export let jpegQualityValue = writable(KEY_JPEG_QUALITY_VALUE, 75);
export let gifColorsValue = writable(KEY_GIF_COLORS_VALUE, 256);
export let tiffCompressionValue = writable(KEY_TIFF_COMPRESSION_VALUE, 0);
export let pngCompressionValue = writable(KEY_PNG_COMPRESSION_VALUE, -1);
export let autoExifOrientation = writable(KEY_EXIF_ORIENTATION_VALUE, false);

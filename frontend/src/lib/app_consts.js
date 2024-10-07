// Copyright (c) 2024 Barat Semet (https://github.com/barats)
// Resizem is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//          http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
// EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
// MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

import { version } from '$app/environment';

//Keys for stores
export const KEY_DOING = 'resizem-' + version + '-doing';
export const KEY_ASK_WHERE_TO_SAVE = 'resizem-' + version + '-ask-where-to-save';
export const KEY_FILES_LIST = 'resizem-' + version + '-files-list';
export const KEY_RESULT_LIST = 'resizem-' + version + '-result-list';
export const KEY_FORMAT_VALUE = 'resizem-' + version + '-format-value';
export const KEY_FILTER_VALUE = 'resizem-' + version + '-filter-value';
export const KEY_WIDTH_VALUE = 'resizem-' + version + '-width-value';
export const KEY_HEIGHT_VALUE = 'resizem-' + version + '-height-value';
export const KEY_CPU_VALUE = 'resizem-' + version + '-cpu-usage-value';
export const KEY_JPEG_QUALITY_VALUE = 'resizem-' + version + '-jpeg-quality-value';
export const KEY_GIF_COLORS_VALUE = 'resizem-' + version + '-gif-colors-value';
export const KEY_TIFF_COMPRESSION_VALUE = 'resizem-' + version + '-tiff-compression-value';
export const KEY_PNG_COMPRESSION_VALUE = 'resizem-' + version + '-png-compression-value';
export const KEY_EXIF_ORIENTATION_VALUE = 'resizem-' + version + '-exif-orientation-value';

//Keys for event
export const EVENT_FILE_DROP = 'resizem-' + version + '-file-drop';
export const EVENT_BEFORE_EXIT = 'resizem-' + version + '-before-exit';
export const EVENT_FILE_RESULT = 'resizem-' + version + '-file-result';
export const EVENT_START = 'resizem-' + version + '-start-jobs';
export const EVENT_CANCEL = 'resizem-' + version + '-cancel-jobs';
export const EVENT_CLEAR_HISTORY = 'resizem-' + version + '-clear-history';
export const EVENT_BACKEND_ERROR = 'resizem-' + version + '_backend_error';

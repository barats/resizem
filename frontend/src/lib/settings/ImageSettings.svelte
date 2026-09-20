<!-- Copyright (c) 2024 Barat Semet (https://github.com/barats)
Resizem is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details. -->

<script>
	import {
		autoExifOrientation,
		gifColorsValue,
		jpegQualityValue,
		pngCompressionValue,
		tiffCompressionValue
	} from '$lib/app_stores';
	import { Label, Range, Select, Toggle } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';

	let tiffCompressionOptions = [
		{ name: 'Uncompressed', value: 0 },
		{ name: 'Deflate', value: 1 },
		{ name: 'LZW', value: 2 },
		{ name: 'CCITTGroup3', value: 3 },
		{ name: 'CCITTGroup4', value: 4 }
	];

	let pngCompressionOptions = [
		{ name: 'Default Compression', value: 0 },
		{ name: 'No Compression', value: -1 },
		{ name: 'Best Speed', value: -2 },
		{ name: 'Best Compression', value: -3 }
	];
</script>

<section>
	<Label class="text-sm font-semibold text-gray-900">{$_('settings.image.title')}</Label>
	<div class="mt-1">
		<div class="grid grid-cols-2 items-center gap-4 border-b border-gray-100 py-4">
			<div>
				<p class="text-sm font-medium text-gray-900">{$_('settings.image.jpg_title')}</p>
				<p class="text-sm text-gray-500">{$_('settings.image.jpg_desc')}</p>
			</div>
			<div style="--wails-draggable:no-drag" class="pl-5">
				<Range
					min="1"
					max="100"
					bind:value={$jpegQualityValue}
					aria-label={$_('settings.image.jpg_title')}
				/>
				<div class="mt-2 text-sm text-gray-500">
					{$_('settings.image.jpg_cur_value')}
					{$jpegQualityValue}
				</div>
			</div>
		</div>
		<div class="grid grid-cols-2 items-center gap-4 border-b border-gray-100 py-4">
			<div>
				<p class="text-sm font-medium text-gray-900">{$_('settings.image.gif_title')}</p>
				<p class="text-sm text-gray-500">{$_('settings.image.gif_desc')}</p>
			</div>
			<div style="--wails-draggable:no-drag" class="pl-5">
				<Range
					min="1"
					max="256"
					bind:value={$gifColorsValue}
					aria-label={$_('settings.image.gif_title')}
				/>
				<div class="mt-2 text-sm text-gray-500">
					{$_('settings.image.gif_cur_value')}
					{$gifColorsValue}
				</div>
			</div>
		</div>
		<div class="grid grid-cols-2 items-center gap-4 border-b border-gray-100 py-4">
			<div>
				<p class="text-sm font-medium text-gray-900">{$_('settings.image.tiff_title')}</p>
				<p class="text-sm text-gray-500">{$_('settings.image.tiff_desc')}</p>
			</div>
			<div style="--wails-draggable:no-drag" class="pl-5">
				<Select
					class="bg-white"
					items={tiffCompressionOptions}
					bind:value={$tiffCompressionValue}
					aria-label={$_('settings.image.tiff_title')}
				/>
			</div>
		</div>
		<div class="grid grid-cols-2 items-center gap-4 border-b border-gray-100 py-4">
			<div>
				<p class="text-sm font-medium text-gray-900">{$_('settings.image.png_title')}</p>
				<p class="text-sm text-gray-500">{$_('settings.image.png_desc')}</p>
			</div>
			<div style="--wails-draggable:no-drag" class="pl-5">
				<Select
					class="bg-white"
					items={pngCompressionOptions}
					bind:value={$pngCompressionValue}
					aria-label={$_('settings.image.png_title')}
				/>
			</div>
		</div>
		<div class="grid grid-cols-2 items-center gap-4 py-4">
			<div>
				<p class="text-sm font-medium text-gray-900">
					{$_('settings.image.exif_orientation_title')}
				</p>
				<p class="text-sm text-gray-500">{$_('settings.image.exif_orientation_desc')}</p>
			</div>
			<div class="pl-5">
				<Toggle bind:checked={$autoExifOrientation}
					>{$_('settings.image.exif_auto_orientation')}</Toggle
				>
			</div>
		</div>
	</div>
</section>

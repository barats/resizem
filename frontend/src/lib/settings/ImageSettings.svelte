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
	import Select from '$lib/ui/Select.svelte';
	import Toggle from '$lib/ui/Toggle.svelte';
	import {
		autoExifOrientation,
		gifColorsValue,
		jpegQualityValue,
		pngCompressionValue,
		tiffCompressionValue
	} from '$lib/app_stores';
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
	<h2 class="text-sm font-semibold text-ink">{$_('settings.image.title')}</h2>
	<div class="mt-1 divide-y divide-line">
		<div class="grid grid-cols-[1fr_16rem] items-center gap-6 py-4 xl:grid-cols-[1fr_20rem]">
			<div>
				<p class="text-sm font-medium text-ink">{$_('settings.image.jpg_title')}</p>
				<p class="text-sm text-ink-muted">{$_('settings.image.jpg_desc')}</p>
			</div>
			<div>
				<input
					type="range"
					min="1"
					max="100"
					bind:value={$jpegQualityValue}
					aria-label={$_('settings.image.jpg_title')}
					class="w-full"
				/>
				<div class="mt-2 text-sm text-ink-secondary">
					{$_('settings.image.jpg_cur_value')}
					{$jpegQualityValue}
				</div>
			</div>
		</div>
		<div class="grid grid-cols-[1fr_16rem] items-center gap-6 py-4 xl:grid-cols-[1fr_20rem]">
			<div>
				<p class="text-sm font-medium text-ink">{$_('settings.image.gif_title')}</p>
				<p class="text-sm text-ink-muted">{$_('settings.image.gif_desc')}</p>
			</div>
			<div>
				<input
					type="range"
					min="1"
					max="256"
					bind:value={$gifColorsValue}
					aria-label={$_('settings.image.gif_title')}
					class="w-full"
				/>
				<div class="mt-2 text-sm text-ink-secondary">
					{$_('settings.image.gif_cur_value')}
					{$gifColorsValue}
				</div>
			</div>
		</div>
		<div class="grid grid-cols-[1fr_16rem] items-center gap-6 py-4 xl:grid-cols-[1fr_20rem]">
			<div>
				<p class="text-sm font-medium text-ink">{$_('settings.image.tiff_title')}</p>
				<p class="text-sm text-ink-muted">{$_('settings.image.tiff_desc')}</p>
			</div>
			<Select
				items={tiffCompressionOptions}
				bind:value={$tiffCompressionValue}
				aria-label={$_('settings.image.tiff_title')}
			/>
		</div>
		<div class="grid grid-cols-[1fr_16rem] items-center gap-6 py-4 xl:grid-cols-[1fr_20rem]">
			<div>
				<p class="text-sm font-medium text-ink">{$_('settings.image.png_title')}</p>
				<p class="text-sm text-ink-muted">{$_('settings.image.png_desc')}</p>
			</div>
			<Select
				items={pngCompressionOptions}
				bind:value={$pngCompressionValue}
				aria-label={$_('settings.image.png_title')}
			/>
		</div>
		<div class="grid grid-cols-[1fr_16rem] items-center gap-6 py-4 xl:grid-cols-[1fr_20rem]">
			<div>
				<p class="text-sm font-medium text-ink">{$_('settings.image.exif_orientation_title')}</p>
				<p class="text-sm text-ink-muted">{$_('settings.image.exif_orientation_desc')}</p>
			</div>
			<Toggle bind:checked={$autoExifOrientation}
				>{$_('settings.image.exif_auto_orientation')}</Toggle
			>
		</div>
	</div>
</section>

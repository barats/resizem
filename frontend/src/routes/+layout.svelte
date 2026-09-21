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
	// @ts-nocheck

	import AppToolbar from '$lib/AppToolbar.svelte';
	import AppStatusBar from '$lib/AppStatusBar.svelte';
	import { onMount } from 'svelte';
	import '../app.css';
	import { X } from '@lucide/svelte';
	import { EventsEmit, EventsOn } from '$lib/wailsjs/runtime/runtime.js';
	import {
		askWheretoSave,
		autoExifOrientation,
		cpuUsageValue,
		doing,
		filesList,
		filterValue,
		formatValue,
		gifColorsValue,
		heightValue,
		jpegQualityValue,
		pngCompressionValue,
		resultList,
		tiffCompressionValue,
		widthValue
	} from '$lib/app_stores';
	import {
		EVENT_BACKEND_ERROR,
		EVENT_BEFORE_EXIT,
		EVENT_CANCEL,
		EVENT_CLEAR_HISTORY,
		EVENT_FILE_DROP,
		EVENT_FILE_RESULT,
		EVENT_START
	} from '$lib/app_consts';
	import { OpenDirectoryDialog, StartHandleFiles } from '$lib/wailsjs/go/rmanager/FileManager';

	let { children } = $props();

	let destPath;
	let err_msg = $state('');
	let showAlert = $state(false);

	onMount(() => {
		EventsOn(EVENT_BACKEND_ERROR, (message) => {
			err_msg = message;
			showAlert = true;
		});

		//On Cancel
		EventsOn(EVENT_CANCEL, () => {
			$doing = false;
		});

		//On Clear History
		EventsOn(EVENT_CLEAR_HISTORY, () => {
			$filesList = [];
			$resultList = [];
		});

		//On Start
		EventsOn(EVENT_START, () => {
			if ($filesList.length <= 0) {
				return;
			}

			if ($askWheretoSave) {
				OpenDirectoryDialog().then((path) => {
					if (path != '' && path != null) {
						destPath = path;
					}
					startJobs();
				});
			} else {
				startJobs();
			} //end of if
		});

		//Backend Event: On File Drop
		EventsOn(EVENT_FILE_DROP, (files) => {
			if (!$doing) {
				$resultList = [];
				$filesList.push(files);
				//Remove duplicate file from array
				$filesList = $filesList.filter((value, index) => $filesList.indexOf(value) === index);
			}
		});

		//Backend Event: On Before Exit
		EventsOn(EVENT_BEFORE_EXIT, () => {
			EventsEmit(EVENT_CANCEL);
			EventsEmit(EVENT_CLEAR_HISTORY);
		});

		//Backend Event: On File Result
		EventsOn(EVENT_FILE_RESULT, (file) => {
			$resultList.push(file);
			$resultList = $resultList;
		});
	});

	function startJobs() {
		var allFiles = $filesList;
		var opts = createImageOptions(destPath);
		$doing = true;
		$filesList = [];
		StartHandleFiles(allFiles, opts)
			.then(() => {})
			.finally(() => {
				$doing = false;
			});
	}

	function createImageOptions(path) {
		return {
			dest_format: $formatValue,
			resample_filter: $filterValue,
			desc_path: path,
			dest_width: $widthValue === null || $widthValue <= 0 ? 0 : $widthValue,
			dest_height: $heightValue === null || $heightValue <= 0 ? 0 : $heightValue,
			jpeg_quality: $jpegQualityValue,
			gif_number_of_colors: $gifColorsValue,
			tiff_compression: $tiffCompressionValue,
			png_compression: $pngCompressionValue,
			auto_orientation: $autoExifOrientation,
			cpu_memory_usage: $cpuUsageValue
		};
	}
</script>

<div class="flex h-dvh flex-col overflow-hidden">
	<AppToolbar />
	<div class="flex min-h-0 flex-1 flex-col overflow-y-auto px-5 py-4">
		{#if showAlert}
			<div
				role="alert"
				class="mb-4 flex shrink-0 items-start gap-2 rounded-lg border border-red-200 bg-red-50 px-3 py-2 text-sm text-red-700"
			>
				<p class="flex-1 break-all">{err_msg}</p>
				<button
					type="button"
					class="shrink-0 rounded-md p-1 text-red-500 transition-colors hover:bg-red-100 hover:text-red-700 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-red-600"
					aria-label="Close"
					onclick={() => (showAlert = false)}
				>
					<X class="h-4 w-4" />
				</button>
			</div>
		{/if}
		{@render children()}
	</div>
	<AppStatusBar />
</div>

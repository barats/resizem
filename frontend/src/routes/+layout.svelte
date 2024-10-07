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

	import AppSidebar from '$lib/AppSidebar.svelte';
	import { onMount } from 'svelte';
	import '../app.css';
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

	let destPath;

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
			$formatValue = 0;
			$filterValue = 1;
			$widthValue = 0;
			$heightValue = 0;
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
					console.log('remote call OpenDirectoryDialog() done');
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
		EventsOn(EVENT_BEFORE_EXIT, (msg) => {
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
			.then(() => {
				console.log('remote call StartHandleFiles() done');
			})
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

<div id="main-content" class="flex overflow-hidden pl-5 pr-5 pt-8">
	<AppSidebar />
	<div id="right-content" class="fixed h-full w-full overflow-hidden pl-20 pr-10">
		<slot />
	</div>
</div>

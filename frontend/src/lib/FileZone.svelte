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
	import { Dropzone } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';
	import { doing, filesList } from '$lib/app_stores';
	import { OpenFilesDialog } from '$lib/wailsjs/go/rmanager/FileManager';
</script>

<Dropzone
	id="dropzone"
	style="--wails-drop-target:drop"
	on:dragover={(event) => {
		event.preventDefault();
	}}
	on:drop={(event) => {
		event.preventDefault();
	}}
	on:dragenter={(event) => {
		event.preventDefault();
	}}
	on:click={(event) => {
		event.preventDefault();
		if (!$doing) {
			OpenFilesDialog().then(() => {
				console.log('remote call OpenFilesDialog() done');
			});
		}
	}}
>
	<svg
		aria-hidden="true"
		class="mb-3 h-10 w-10 text-gray-400"
		fill="none"
		stroke="currentColor"
		viewBox="0 0 24 24"
		xmlns="http://www.w3.org/2000/svg"
		><path
			stroke-linecap="round"
			stroke-linejoin="round"
			stroke-width="2"
			d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"
		/></svg
	>
	<p class="mb-2 text-lg text-gray-500">
		{@html $_('home.dropzone.message')}
	</p>
	<p class="mb-2 text-xs text-gray-500">
		{$_('home.dropzone.types')}
	</p>
	{#if $filesList.length > 0}
		<p>
			{$_('home.dropzone.files_selected', { values: { count: $filesList.length } })}
		</p>
	{/if}
</Dropzone>

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
	import { CloudUpload } from '@lucide/svelte';
	import { _ } from 'svelte-i18n';
	import { doing, filesList } from '$lib/app_stores';
	import { OpenFilesDialog } from '$lib/wailsjs/go/rmanager/FileManager';

	function openDialog() {
		if (!$doing) {
			OpenFilesDialog().then(() => {});
		}
	}
</script>

<!-- The wrapper marks this region as the Wails file-drop target. -->
<div style="--wails-drop-target:drop" class="h-full">
	<!-- A real button: keyboard users get visible focus and Enter/Space
		activation instead of the old sr-only fallback link. -->
	<button
		type="button"
		disabled={$doing}
		class="flex h-full w-full flex-col items-center justify-center gap-1.5 rounded-xl border border-line bg-surface px-6 text-center transition-colors hover:border-primary-400 hover:bg-primary-50/50 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600 disabled:pointer-events-none disabled:opacity-60"
		onclick={openDialog}
	>
		<CloudUpload class="mb-1 h-10 w-10 text-ink-muted" strokeWidth={1.5} />
		{#if $doing}
			<p class="text-sm font-medium text-ink-secondary">
				{$_('home.dropzone.doing')}
			</p>
		{:else}
			<p class="text-base font-semibold text-ink">
				{$_('home.dropzone.upload')}
			</p>
		{/if}
		<p class="text-xs text-ink-muted">
			{$_('home.dropzone.types')}
		</p>
		{#if $filesList.length > 0}
			<p class="text-xs text-ink-secondary">
				{$_('home.dropzone.files_selected', { values: { count: $filesList.length } })}
			</p>
		{/if}
	</button>
</div>

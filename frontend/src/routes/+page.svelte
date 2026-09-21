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
	import AppFilesList from '$lib/AppFilesList.svelte';
	import AppOptions from '$lib/AppOptions.svelte';
	import FileZone from '$lib/FileZone.svelte';

	let scrollableDiv = $state();

	$effect(() => {
		if (scrollableDiv) {
			scrollToBottom(scrollableDiv);
		}
	});

	function scrollToBottom(node) {
		node.scroll({ top: node.scrollHeight, behavior: 'smooth' });
	}
</script>

<main class="flex h-full min-h-0 w-full flex-col gap-4">
	<!-- Row 1: dropzone and options at 2:1. The options card sizes the row
		and the dropzone stretches to match it. -->
	<div class="grid shrink-0 grid-cols-[2fr_1fr] gap-4">
		<FileZone />
		<AppOptions />
	</div>
	<!-- Row 2: queue and results -->
	<div bind:this={scrollableDiv} class="min-h-24 flex-1 overflow-y-auto">
		<AppFilesList />
	</div>
</main>

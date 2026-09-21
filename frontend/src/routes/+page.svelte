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

<!-- Two rows at 1:2. Row 1 (dropzone + options at 2:1) floors at its
	content height when the window is short; the ratio holds whenever
	the options card fits in the 1fr share. -->
<main class="grid h-full min-h-0 w-full grid-rows-[1fr_2fr] gap-4">
	<div class="grid grid-cols-[2fr_1fr] gap-4">
		<FileZone />
		<AppOptions />
	</div>
	<!-- Row 2: one dashed collection panel holding queue and results -->
	<div
		bind:this={scrollableDiv}
		class="overflow-y-auto rounded-xl border border-dashed border-line p-2"
	>
		<AppFilesList />
	</div>
</main>

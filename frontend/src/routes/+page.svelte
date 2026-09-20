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
	import { filesList, resultList } from '$lib/app_stores';
	import AppButtons from '$lib/AppButtons.svelte';

	let scrollableDiv = $state();

	$effect(() => {
		if (scrollableDiv && ($filesList.length > 0 || $resultList.length > 0)) {
			scrollToBottom(scrollableDiv);
		}
	});

	function scrollToBottom(node) {
		node.scroll({ top: node.scrollHeight, behavior: 'smooth' });
	}
</script>

<main class="flex h-full min-h-0 w-full flex-col gap-5">
	<div class="grid h-max shrink-0 grid-cols-3 gap-5">
		<div class="col-span-2">
			<FileZone />
		</div>
		<div>
			<AppOptions />
		</div>
	</div>
	<div
		bind:this={scrollableDiv}
		class="min-h-0 flex-1 overflow-y-auto rounded-xl border border-gray-100 bg-white"
	>
		<AppFilesList />
	</div>
	<div class="flex shrink-0 items-center justify-end gap-3">
		<AppButtons />
	</div>
</main>

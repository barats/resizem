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
	import { Listgroup, ListgroupItem } from 'flowbite-svelte';
	import {
		FileImageOutline,
		CheckCircleOutline,
		CloseCircleOutline,
		CloseOutline
	} from 'flowbite-svelte-icons';
	import { filesList, resultList } from '$lib/app_stores';
	import { _ } from 'svelte-i18n';

	let queueCount = $derived($filesList.length);
	let resultCount = $derived($resultList.length);
	let isEmpty = $derived(queueCount === 0 && resultCount === 0);

	let friendlyErrors = $derived([
		{ pattern: /unknown format|unsupported/i, text: $_('home.errors.unknown_format') },
		{ pattern: /eof|truncat/i, text: $_('home.errors.unexpected_eof') }
	]);

	function friendlyMessage(message) {
		if (!message) {
			return '';
		}
		const match = friendlyErrors.find((entry) => entry.pattern.test(message));
		return match ? match.text : message;
	}

	function removeQueuedFile(path) {
		$filesList = $filesList.filter((file) => file !== path);
	}
</script>

{#if isEmpty}
	<div class="flex h-full flex-col items-center justify-center gap-2 py-10 text-center">
		<FileImageOutline class="h-10 w-10 text-gray-300" />
		<p class="text-sm text-gray-400">{$_('home.list.empty')}</p>
	</div>
{:else}
	<div class="flex flex-col gap-5 pb-3">
		{#if queueCount > 0}
			<section>
				<h2 class="mb-1 px-1 text-sm font-medium text-gray-500">
					{$_('home.list.queue')} ({queueCount})
				</h2>
				<Listgroup class="w-full">
					{#each $filesList as file (file)}
						<ListgroupItem class="text-sm">
							<FileImageOutline class="h-5 w-5 shrink-0 text-gray-400" />
							<span class="min-w-0 flex-1 truncate" title={file}>{file}</span>
							<button
								type="button"
								class="shrink-0 rounded p-1 text-gray-400 hover:bg-gray-100 hover:text-gray-700 focus:ring-2 focus:ring-gray-200 focus:outline-none"
								aria-label={$_('home.list.remove')}
								title={$_('home.list.remove')}
								onclick={() => removeQueuedFile(file)}
							>
								<CloseOutline class="h-4 w-4" />
							</button>
						</ListgroupItem>
					{/each}
				</Listgroup>
			</section>
		{/if}

		{#if resultCount > 0}
			<section>
				<h2 class="mb-1 px-1 text-sm font-medium text-gray-500">
					{$_('home.list.results')} ({resultCount})
				</h2>
				<Listgroup class="w-full">
					{#each $resultList as item, i (i)}
						<ListgroupItem class="text-sm">
							{#if item.status === 1}
								<CheckCircleOutline class="h-5 w-5 shrink-0 text-green-600" />
								<span class="min-w-0 flex-1 truncate" title={item.name}>{item.name}</span>
							{:else}
								<CloseCircleOutline class="h-5 w-5 shrink-0 text-red-600" />
								<span class="min-w-0 flex-1 truncate" title={item.name}>{item.name}</span>
								<span class="shrink-0 text-xs text-red-600">{friendlyMessage(item.message)}</span>
							{/if}
						</ListgroupItem>
					{/each}
				</Listgroup>
			</section>
		{/if}
	</div>
{/if}

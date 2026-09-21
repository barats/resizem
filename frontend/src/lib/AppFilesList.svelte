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
	import Button from '$lib/ui/Button.svelte';
	import { CircleCheck, CircleX, FileImage, X } from '@lucide/svelte';
	import { doing, filesList, resultList } from '$lib/app_stores';
	import { _ } from 'svelte-i18n';
	import { EventsEmit } from './wailsjs/runtime/runtime';
	import { EVENT_CANCEL, EVENT_START } from './app_consts';
	import { CancelHandleFiles } from './wailsjs/go/rmanager/FileManager';

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

	function startJobs() {
		EventsEmit(EVENT_START);
	}

	function cancelJobs() {
		CancelHandleFiles().then(() => {});
		EventsEmit(EVENT_CANCEL);
	}

	function clearQueue() {
		$filesList = [];
	}

	function clearResults() {
		$resultList = [];
	}
</script>

{#if isEmpty}
	<div class="flex h-full flex-col items-center justify-center gap-2 py-10 text-center">
		<FileImage class="h-10 w-10 text-zinc-300" strokeWidth={1.5} />
		<p class="text-sm text-ink-muted">
			{#if $doing}
				{$_('home.dropzone.doing')}
			{:else}
				{$_('home.list.empty')}
			{/if}
		</p>
	</div>
{:else}
	<div class="flex flex-col gap-5 pb-1">
		{#if queueCount > 0}
			<section>
				<div class="mb-1.5 flex items-center justify-between px-1">
					<h2 class="text-sm font-medium text-ink-secondary">
						{$_('home.list.queue')} ({queueCount})
					</h2>
					{#if !$doing}
						<div class="flex items-center gap-2">
							<Button variant="primary" size="sm" onclick={startJobs}>
								{$_('home.buttons.start')}
							</Button>
							<Button variant="ghost" size="sm" onclick={clearQueue}>
								{$_('home.list.clear_queue')}
							</Button>
						</div>
					{/if}
				</div>
				<ul class="divide-y divide-line">
					{#each $filesList as file (file)}
						<li
							class="group flex cursor-pointer items-center gap-3 px-3 py-2 text-sm font-normal text-ink transition-colors hover:bg-surface-muted"
						>
							<FileImage class="h-4.5 w-4.5 shrink-0 text-ink-muted" />
							<span class="min-w-0 flex-1 truncate" title={file}>{file}</span>
							{#if !$doing}
								<button
									type="button"
									class="shrink-0 rounded-md p-1 text-ink-muted opacity-0 transition-opacity group-hover:opacity-100 hover:bg-line hover:text-ink focus-visible:opacity-100 focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
									aria-label={$_('home.list.remove')}
									title={$_('home.list.remove')}
									onclick={() => removeQueuedFile(file)}
								>
									<X class="h-4 w-4" />
								</button>
							{/if}
						</li>
					{/each}
				</ul>
			</section>
		{/if}

		{#if resultCount > 0 || $doing}
			<section>
				<div class="mb-1.5 flex items-center justify-between px-1">
					<h2 class="text-sm font-medium text-ink-secondary">
						{$_('home.list.results')} ({resultCount})
					</h2>
					{#if $doing}
						<div class="flex items-center gap-3">
							<span class="text-sm text-ink-secondary">
								{$_('home.list.converting', { values: { count: resultCount } })}
							</span>
							<Button variant="danger" size="sm" onclick={cancelJobs}>
								{$_('home.buttons.cancel')}
							</Button>
						</div>
					{:else}
						<Button variant="ghost" size="sm" onclick={clearResults}>
							{$_('home.list.clear_results')}
						</Button>
					{/if}
				</div>
				<ul class="divide-y divide-line">
					{#each $resultList as item, i (i)}
						<li
							class="flex cursor-pointer items-center gap-3 px-3 py-2 text-sm font-normal text-ink transition-colors hover:bg-surface-muted"
						>
							{#if item.status === 1}
								<CircleCheck class="h-4.5 w-4.5 shrink-0 text-green-600" />
								<span class="min-w-0 flex-1 truncate" title={item.name}>{item.name}</span>
							{:else}
								<CircleX class="h-4.5 w-4.5 shrink-0 text-red-600" />
								<span class="min-w-0 flex-1 truncate" title={item.name}>{item.name}</span>
								<span class="shrink-0 text-xs text-red-600">{friendlyMessage(item.message)}</span>
							{/if}
						</li>
					{/each}
				</ul>
			</section>
		{/if}
	</div>
{/if}

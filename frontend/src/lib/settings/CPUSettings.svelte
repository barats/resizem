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
	import { cpuUsageValue } from '$lib/app_stores';
	import { Label, Range } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';

	// Slider step is 1; round defensively in case an older persisted value is fractional.
	let cpuLevel = $derived(Math.min(3, Math.max(1, Math.round($cpuUsageValue || 1))));
</script>

<section>
	<Label for="range-minmax" class="text-sm font-semibold text-gray-900"
		>{$_('settings.cpu.title')}</Label
	>
	<div class="mt-3 grid grid-cols-2 items-center gap-5">
		<div style="--wails-draggable:no-drag">
			<Range
				id="range-minmax"
				min="1"
				max="3"
				bind:value={$cpuUsageValue}
				aria-label={$_('settings.cpu.title')}
			/>
			<div class="grid grid-cols-3 pt-2">
				<span class="text-sm {cpuLevel === 1 ? 'font-medium text-gray-900' : 'text-gray-500'}"
					>{$_('settings.cpu.medium')}</span
				>
				<span
					class="text-center text-sm {cpuLevel === 2
						? 'font-medium text-gray-900'
						: 'text-gray-500'}">{$_('settings.cpu.high')}</span
				>
				<span
					class="text-right text-sm {cpuLevel === 3
						? 'font-medium text-gray-900'
						: 'text-gray-500'}">{$_('settings.cpu.most')}</span
				>
			</div>
		</div>
		<div class="pl-5">
			<p class="text-sm text-gray-500">
				{#if cpuLevel === 1}
					{$_('settings.cpu.medium_desc')}
				{:else if cpuLevel === 2}
					{$_('settings.cpu.high_desc')}
				{:else if cpuLevel === 3}
					{$_('settings.cpu.most_desc')}
				{/if}
			</p>
		</div>
	</div>
</section>

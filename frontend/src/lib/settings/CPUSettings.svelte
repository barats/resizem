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
</script>

<Label for="range-minmax" defaultClass="text-lg font-medium">{$_('settings.cpu.title')}</Label>
<div class="grid grid-cols-2 gap-5">
	<div style="--wails-draggable:no-drag">
		<Range id="range-minmax" min="1" max="3" bind:value={$cpuUsageValue} />
		<div class="grid grid-cols-3 pt-5">
			<span class="text-sm text-gray-500">{$_('settings.cpu.medium')}</span>
			<span class="text-center text-sm text-gray-500">{$_('settings.cpu.high')}</span>
			<span class="text-right text-sm text-gray-500">{$_('settings.cpu.most')}</span>
		</div>
	</div>
	<div class="pl-5">
		{#if $cpuUsageValue === 1}
			<p>{$_('settings.cpu.medium_desc')}</p>
		{:else if $cpuUsageValue === 1.5}
			<p>{$_('settings.cpu.high_desc')}</p>
		{:else if $cpuUsageValue === 2}
			<p>{$_('settings.cpu.most_desc')}</p>
		{/if}
	</div>
</div>

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
	import Input from '$lib/ui/Input.svelte';
	import Select from '$lib/ui/Select.svelte';
	import { OutputImagesTypes, ResampleFilterTypes } from '$lib/wailsjs/go/rmanager/TypeManager.js';
	import { onMount } from 'svelte';
	import { _ } from 'svelte-i18n';
	import { filterValue, formatValue, heightValue, widthValue } from './app_stores';

	let allFormats = $state();
	let allFilters = $state();

	onMount(() => {
		OutputImagesTypes().then((data) => {
			var keep = { name: $_('home.options.keep'), value: 0 };
			data.push(keep);
			allFormats = data;
		});

		ResampleFilterTypes().then((data) => {
			allFilters = data;
		});
	});

	// null means the field is empty ("not set"), which is valid; only negatives are invalid.
	let showWidthHelper = $derived($widthValue < 0);
	let showHeightHelper = $derived($heightValue < 0);
</script>

<div class="grid content-start gap-3 rounded-xl border border-line bg-surface p-3">
	<div class="grid grid-cols-2 gap-3">
		<div class="grid gap-1.5">
			<label for="format-select" class="text-xs font-medium text-ink-secondary"
				>{$_('home.options.format')}</label
			>
			<Select
				id="format-select"
				placeholder={$_('home.options.choose')}
				items={allFormats}
				bind:value={$formatValue}
			/>
		</div>
		<div class="grid gap-1.5">
			<label for="filter-select" class="text-xs font-medium text-ink-secondary"
				>{$_('home.options.filter')}</label
			>
			<Select id="filter-select" items={allFilters} bind:value={$filterValue} />
		</div>
	</div>
	<div class="grid grid-cols-2 gap-3">
		<div class="grid gap-1.5">
			<label for="width" class="text-xs font-medium text-ink-secondary"
				>{$_('home.options.width.title')}</label
			>
			<Input id="width" type="number" bind:value={$widthValue} />
		</div>
		<div class="grid gap-1.5">
			<label for="height" class="text-xs font-medium text-ink-secondary"
				>{$_('home.options.height.title')}</label
			>
			<Input id="height" type="number" bind:value={$heightValue} />
		</div>
		{#if showWidthHelper || showHeightHelper}
			<p class="col-span-2 text-xs font-medium text-red-600">
				{$_('home.options.width.helper1')}
				{$_('home.options.width.helper2')}
			</p>
		{:else}
			<p class="col-span-2 text-xs text-ink-muted">{$_('home.options.size_helper')}</p>
		{/if}
	</div>
</div>

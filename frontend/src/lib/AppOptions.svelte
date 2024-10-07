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
	import { NumberInput, Label, Select, Helper } from 'flowbite-svelte';
	import { OutputImagesTypes, ResampleFilterTypes } from '$lib/wailsjs/go/rmanager/TypeManager.js';
	import { onMount } from 'svelte';
	import { _ } from 'svelte-i18n';
	import { filterValue, formatValue, heightValue, widthValue } from './app_stores';

	let showWidthHelper = false,
		showHeightHelper = false;

	let allFormats, allFilters;

	onMount(() => {
		OutputImagesTypes().then((data) => {
			var keep = { name: $_('home.options.keep'), value: 0 };
			data.push(keep);
			allFormats = data;
			console.log('remote call OutputImagesTypes() done');
		});

		ResampleFilterTypes().then((data) => {
			allFilters = data;
			console.log('remote call ResampleFilterTypes() done');
		});
	});

	$: if ($widthValue < 0 || $widthValue === null) {
		showWidthHelper = true;
	} else {
		showWidthHelper = false;
	}

	$: if ($heightValue < 0 || $heightValue === null) {
		showHeightHelper = true;
	} else {
		showHeightHelper = false;
	}
</script>

<div class="grid grid-cols-1 gap-2">
	<div>
		<Label for="format-select">{$_('home.options.format')}</Label>
		<Select
			id="format-select"
			placeholder={$_('home.options.choose')}
			items={allFormats}
			bind:value={$formatValue}
		/>
	</div>
	<div>
		<Label for="filter-select">{$_('home.options.filter')}</Label>
		<Select id="filter-select" items={allFilters} bind:value={$filterValue} />
	</div>
	<div class="grid grid-cols-2 gap-5">
		<Label
			>{$_('home.options.width.title')}
			<NumberInput id="width" bind:value={$widthValue} />
			{#if showWidthHelper}
				<Helper class="mt-2" color="red">
					<span class="font-medium">{$_('home.options.width.helper1')}</span>
					{$_('home.options.width.helper2')}
				</Helper>
			{/if}
		</Label>
		<Label
			>{$_('home.options.height.title')}
			<NumberInput id="height" bind:value={$heightValue} />
			{#if showHeightHelper}
				<Helper class="mt-2" color="red">
					<span class="font-medium">{$_('home.options.height.helper1')}</span>
					{$_('home.options.height.helper2')}
				</Helper>
			{/if}
		</Label>
	</div>
</div>

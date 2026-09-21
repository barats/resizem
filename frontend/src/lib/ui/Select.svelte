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
	import { ChevronDown } from '@lucide/svelte';

	let {
		id = undefined,
		items = [],
		placeholder = '',
		class: klass = '',
		value = $bindable(),
		...rest
	} = $props();

	let selectEl = $state(null);

	// Options usually arrive async (Go backend lists), after the initial
	// binding run, so re-apply the bound value once they mount.
	$effect(() => {
		if (selectEl && items.length > 0 && value !== undefined && value !== null) {
			selectEl.value = String(value);
		}
	});
</script>

<div class="relative">
	<select
		bind:this={selectEl}
		{id}
		bind:value
		class="h-9 w-full appearance-none rounded-lg border border-line bg-surface pr-9 pl-3 text-sm text-ink transition-colors hover:border-line-strong focus:border-primary-500 focus:ring-2 focus:ring-primary-500/25 focus:outline-none {klass}"
		{...rest}
	>
		{#if placeholder}
			<option value="" disabled hidden>{placeholder}</option>
		{/if}
		{#each items as item (item.value)}
			<option value={item.value}>{item.name}</option>
		{/each}
	</select>
	<ChevronDown
		aria-hidden="true"
		class="pointer-events-none absolute top-1/2 right-3 h-4 w-4 -translate-y-1/2 text-ink-muted"
	/>
</div>

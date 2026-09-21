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
	let { checked = $bindable(false), class: klass = '', children, ...rest } = $props();
</script>

<label class="inline-flex cursor-pointer items-center gap-2 {klass}">
	<input type="checkbox" bind:checked class="sr-only" {...rest} />
	<span
		aria-hidden="true"
		class="relative h-5 w-9 shrink-0 rounded-full transition-colors has-focus-visible:ring-2 has-focus-visible:ring-primary-500 has-focus-visible:ring-offset-2 {checked
			? 'bg-primary-600'
			: 'bg-zinc-200'}"
	>
		<span
			class="absolute top-0.5 left-0.5 h-4 w-4 rounded-full bg-white shadow-sm transition-transform {checked
				? 'translate-x-4'
				: ''}"
		></span>
	</span>
	{@render children?.()}
</label>

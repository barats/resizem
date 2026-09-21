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
	let { id = undefined, type = 'text', class: klass = '', value = $bindable(), ...rest } = $props();
</script>

<input
	{id}
	{type}
	bind:value
	class="h-9 w-full rounded-lg border border-line bg-surface px-3 text-sm text-ink transition-colors placeholder:text-ink-muted hover:border-line-strong focus:border-primary-500 focus:ring-2 focus:ring-primary-500/25 focus:outline-none {klass}"
	{...rest}
/>

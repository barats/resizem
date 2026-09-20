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
	import { Button, Modal } from 'flowbite-svelte';
	import { _, json } from 'svelte-i18n';

	let { visible = $bindable(false) } = $props();

	let aboutContent = $derived(/** @type {string[]} */ ($json('about.content')) ?? []);
</script>

<Modal bind:open={visible} title={$_('about.title')}>
	{#each aboutContent as item, i (i)}
		<!-- eslint-disable-next-line svelte/no-at-html-tags -- static legal copy bundled with the app, not user input -->
		{@html item}
	{/each}
	{#snippet footer()}
		<div class="w-full text-right">
			<Button onclick={() => (visible = false)}>{$_('about.okay_btn')}</Button>
		</div>
	{/snippet}
</Modal>

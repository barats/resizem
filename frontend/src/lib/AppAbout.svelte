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
	// @ts-nocheck
	import { Button, Modal } from 'flowbite-svelte';
	import { _, json } from 'svelte-i18n';
	export let visible = false;
	export let onClose;
</script>

<Modal
	title={$_('about.title')}
	bind:open={visible}
	on:close={() => {
		onClose();
	}}
>
	{#each $json('about.content') as item}
		{@html item}
	{/each}
	<svelte:fragment slot="footer">
		<div class="w-full text-right">
			<Button on:click={() => (visible = false)}>{$_('about.okay_btn')}</Button>
		</div>
	</svelte:fragment>
</Modal>

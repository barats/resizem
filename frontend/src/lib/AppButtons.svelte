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
	import { Button } from 'flowbite-svelte';
	import { EventsEmit } from './wailsjs/runtime/runtime';
	import { EVENT_CANCEL, EVENT_CLEAR_HISTORY, EVENT_START } from './app_consts';
	import { doing, filesList, resultList } from './app_stores';
	import { _ } from 'svelte-i18n';
	import { CancelHandleFiles } from './wailsjs/go/rmanager/FileManager';

	const handleCancel = () => {
		CancelHandleFiles().then(() => {});
		EventsEmit(EVENT_CANCEL);
	};

	const handleClear = () => {
		EventsEmit(EVENT_CLEAR_HISTORY);
	};

	const handleStart = () => {
		EventsEmit(EVENT_START);
	};
</script>

{#if $doing}
	<Button size="lg" color="blue" class="bg-blue-600" disabled loading>
		{$_('home.buttons.doing')}
	</Button>
	<Button size="lg" color="red" onclick={handleCancel}>
		{$_('home.buttons.cancel')}
	</Button>
{:else}
	<Button
		size="lg"
		color="blue"
		class="bg-blue-600 hover:bg-blue-700"
		disabled={$filesList.length === 0}
		onclick={handleStart}
	>
		{$_('home.buttons.start')}
	</Button>
	<Button
		size="lg"
		color="light"
		disabled={$filesList.length === 0 && $resultList.length === 0}
		onclick={handleClear}
	>
		{$_('home.buttons.clear')}
	</Button>
{/if}

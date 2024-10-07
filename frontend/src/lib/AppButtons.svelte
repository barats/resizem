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
	import { Button, Spinner } from 'flowbite-svelte';
	import { EventsEmit } from './wailsjs/runtime/runtime';
	import { EVENT_CANCEL, EVENT_CLEAR_HISTORY, EVENT_START } from './app_consts';
	import { doing } from './app_stores';
	import { _ } from 'svelte-i18n';
	import { CancelHandleFiles } from './wailsjs/go/rmanager/FileManager';

	const handleCancel = () => {
		CancelHandleFiles().then(() => {
			console.log('remote call CancelHandleFiles() done');
		});
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
	<Button size="xl" shadow color="blue" disabled>
		<Spinner color="white" />{$_('home.buttons.doing')}
	</Button>
	<Button size="xl" shadow color="red" on:click={handleCancel}>{$_('home.buttons.cancel')}</Button>
{:else}
	<Button size="xl" shadow color="blue" on:click={handleStart}>{$_('home.buttons.start')}</Button>
	<Button size="xl" shadow color="red" on:click={handleClear}>{$_('home.buttons.clear')}</Button>
{/if}

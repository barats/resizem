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
	import { page } from '$app/state';
	import { Sidebar, SidebarGroup, SidebarItem, Button } from 'flowbite-svelte';
	import { CogOutline, GithubSolid, ImageSolid, InfoCircleSolid } from 'flowbite-svelte-icons';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';
	import { _ } from 'svelte-i18n';
	import AppAbout from './AppAbout.svelte';

	let visible = $state(false);

	const activeClass =
		'flex items-center justify-center mx-2 p-4 text-base font-normal text-blue-700 bg-blue-50 rounded-lg hover:bg-blue-100';
	const nonActiveClass =
		'flex items-center justify-center mx-2 p-4 text-base font-normal text-gray-500 rounded-lg hover:bg-gray-100 hover:text-gray-900';

	const GithubHomePage = () => {
		BrowserOpenURL('https://github.com/barats/resizem');
	};
</script>

<Sidebar
	activeUrl={page.url.pathname}
	ariaLabel="sidebar"
	position="fixed"
	alwaysOpen
	disableBreakpoints
	class="relative z-20 flex h-full w-20 flex-shrink-0 flex-col"
	classes={{
		div: 'h-full border-r border-gray-100 bg-white px-0 py-0',
		active: activeClass,
		nonactive: nonActiveClass
	}}
>
	<SidebarGroup>
		<SidebarItem href="/" spanClass="hidden" aria-label={$_('nav.home')} title={$_('nav.home')}>
			{#snippet icon()}
				<ImageSolid class="h-5 w-5 transition duration-75" />
			{/snippet}
		</SidebarItem>
		<SidebarItem
			href="/settings/"
			spanClass="hidden"
			aria-label={$_('nav.settings')}
			title={$_('nav.settings')}
		>
			{#snippet icon()}
				<CogOutline class="h-5 w-5 transition duration-75" />
			{/snippet}
		</SidebarItem>
	</SidebarGroup>

	<SidebarGroup
		class="absolute bottom-5 flex w-full flex-col flex-wrap items-center justify-center gap-1"
	>
		<Button
			color="light"
			size="xs"
			class="w-11/12 border-0 p-2!"
			aria-label={$_('nav.github')}
			title={$_('nav.github')}
			onclick={GithubHomePage}
		>
			<GithubSolid class="h-5 w-5" />
		</Button>
		<Button
			color="light"
			size="xs"
			class="w-11/12 border-0 p-2!"
			aria-label={$_('nav.about')}
			title={$_('nav.about')}
			onclick={() => (visible = true)}
		>
			<InfoCircleSolid class="h-5 w-5" />
		</Button>
	</SidebarGroup>
</Sidebar>

<AppAbout bind:visible />

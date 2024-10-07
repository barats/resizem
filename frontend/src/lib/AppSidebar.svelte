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
	import { page } from '$app/stores';
	import { Sidebar, SidebarGroup, SidebarItem, SidebarWrapper, Button } from 'flowbite-svelte';
	import { CogOutline, GithubSolid, ImageSolid, InfoCircleSolid } from 'flowbite-svelte-icons';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';
	import AppAbout from './AppAbout.svelte';

	let visible = false;
	$: activeUrl = $page.url.pathname;

	let activeClass =
		'flex items-center justify-center p-4 text-base font-normal text-primary-900 bg-primary-200 rounded-lg hover:bg-primary-100';
	let nonActiveClass =
		'flex items-center justify-center p-4 text-base font-normal text-primary-900 rounded-lg hover:bg-primary-100';

	const GithubHomePage = () => {
		BrowserOpenURL('https://github.com/barats/resizem');
	};
</script>

<Sidebar
	{activeUrl}
	{activeClass}
	{nonActiveClass}
	ariaLabel="sidebar"
	asideClass="fixed top-0 left-0 pt-10 z-20 flex flex-col flex-shrink-0 w-20 h-full transition-width"
>
	<SidebarWrapper divClass="bg-white h-screen border-r-2">
		<SidebarGroup>
			<SidebarItem href="/" spanClass="display:none;">
				<svelte:fragment slot="icon">
					<ImageSolid
						class="h-5 w-5  text-gray-500 transition duration-75  group-hover:text-gray-900"
					/>
				</svelte:fragment>
			</SidebarItem>
			<SidebarItem href="/settings/" spanClass="display:none;">
				<svelte:fragment slot="icon">
					<CogOutline
						class="h-5 w-5 text-gray-500 transition duration-75  group-hover:text-gray-900"
					/>
				</svelte:fragment>
			</SidebarItem>
		</SidebarGroup>

		<SidebarGroup>
			<div
				class="absolute bottom-5 flex w-full flex-col flex-wrap items-center justify-center gap-1"
			>
				<Button color="light" size="xs" class="w-11/12 border-0 !p-2" on:click={GithubHomePage}>
					<GithubSolid class="h-5 w-5" />
				</Button>
				<Button
					color="light"
					size="xs"
					class="w-11/12 border-0 !p-2"
					on:click={() => {
						visible = true;
					}}
				>
					<InfoCircleSolid class="h-5 w-5" />
				</Button>
			</div>
		</SidebarGroup>
	</SidebarWrapper>
</Sidebar>

<AppAbout
	{visible}
	onClose={() => {
		visible = false;
	}}
/>

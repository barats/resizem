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
	import { resolve } from '$app/paths';
	import { Image, Info, Settings } from '@lucide/svelte';
	import { BrowserOpenURL } from '$lib/wailsjs/runtime/runtime';
	import { _ } from 'svelte-i18n';
	import GithubMark from '$lib/ui/GithubMark.svelte';
	import AppAbout from './AppAbout.svelte';

	let visible = $state(false);

	const GithubHomePage = () => {
		BrowserOpenURL('https://github.com/barats/resizem');
	};

	const navLinkClass = (active) =>
		`flex items-center gap-2 rounded-lg px-3 py-1.5 text-sm font-medium transition-colors focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600 ${
			active
				? 'bg-surface-muted text-ink'
				: 'text-ink-secondary hover:bg-surface-muted hover:text-ink'
		}`;
</script>

<!-- pl-20 keeps the toolbar's left edge clear for the macOS traffic lights
	(FullSizeContent renders content under the title bar). -->
<header
	class="drag-region flex h-12 shrink-0 items-center gap-1 border-b border-line bg-surface pr-3 pl-20"
>
	<nav class="flex items-center gap-1">
		<a
			href={resolve('/')}
			aria-current={page.url.pathname === '/' ? 'page' : undefined}
			title={$_('nav.home')}
			class={navLinkClass(page.url.pathname === '/')}
		>
			<Image class="h-4.5 w-4.5" strokeWidth={2} />
			{$_('nav.home')}
		</a>
		<a
			href={resolve('/settings/')}
			aria-current={page.url.pathname.startsWith('/settings') ? 'page' : undefined}
			title={$_('nav.settings')}
			class={navLinkClass(page.url.pathname.startsWith('/settings'))}
		>
			<Settings class="h-4.5 w-4.5" strokeWidth={2} />
			{$_('nav.settings')}
		</a>
	</nav>

	<div class="ml-auto flex items-center gap-1">
		<button
			type="button"
			class="flex h-9 w-9 items-center justify-center rounded-lg text-ink-secondary transition-colors hover:bg-surface-muted hover:text-ink focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
			aria-label={$_('nav.github')}
			title={$_('nav.github')}
			onclick={GithubHomePage}
		>
			<GithubMark class="h-5 w-5" />
		</button>
		<button
			type="button"
			class="flex h-9 w-9 items-center justify-center rounded-lg text-ink-secondary transition-colors hover:bg-surface-muted hover:text-ink focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-primary-600"
			aria-label={$_('nav.about')}
			title={$_('nav.about')}
			onclick={() => (visible = true)}
		>
			<Info class="h-5 w-5" />
		</button>
	</div>
</header>

<AppAbout bind:visible />

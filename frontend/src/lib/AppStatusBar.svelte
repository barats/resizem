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
	import { version } from '$app/environment';
	import { cpuUsageValue } from './app_stores';
	import { _ } from 'svelte-i18n';

	// Round defensively in case an older persisted value is fractional.
	let cpuLevel = $derived(Math.min(3, Math.max(1, Math.round($cpuUsageValue || 1))));
	let cpuLabel = $derived(
		cpuLevel === 1
			? $_('settings.cpu.medium')
			: cpuLevel === 2
				? $_('settings.cpu.high')
				: $_('settings.cpu.most')
	);
</script>

<footer
	class="flex h-8 shrink-0 items-center justify-between border-t border-line bg-surface px-5 text-xs text-ink-muted"
>
	<span>Resizem {version}</span>
	<span>
		{$_('statusbar.cpu')}:
		<span class="font-medium text-ink-secondary">{cpuLabel}</span>
	</span>
</footer>

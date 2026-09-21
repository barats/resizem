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
	import Select from '$lib/ui/Select.svelte';
	import { _ } from 'svelte-i18n';
	import { appLocales, getAppLocale, setAppLocale } from '../i18n';

	let selected = $state(getAppLocale());
</script>

<section>
	<label for="lang-select" class="text-sm font-semibold text-ink"
		>{$_('settings.language.title')}</label
	>
	<div class="mt-3">
		<Select
			id="lang-select"
			placeholder={$_('home.options.choose')}
			bind:value={selected}
			items={appLocales}
			onchange={() => setAppLocale(selected)}
		></Select>
	</div>
</section>

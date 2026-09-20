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
	import { Label, Select } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';
	import { appLocales, getAppLocale, setAppLocale } from '../i18n';

	let selected = $state(getAppLocale());
</script>

<section>
	<Label for="lang-select" class="text-sm font-semibold text-gray-900"
		>{$_('settings.language.title')}</Label
	>
	<div style="--wails-draggable:no-drag" class="mt-3">
		<Select
			id="lang-select"
			class="bg-white"
			placeholder={$_('home.options.choose')}
			bind:value={selected}
			items={appLocales}
			onchange={() => setAppLocale(selected)}
		></Select>
	</div>
</section>

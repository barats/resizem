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
	import { askWheretoSave } from '$lib/app_stores';
	import { Label, Toggle } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';
</script>

<Label defaultClass="text-lg font-medium" for="lang-select">{$_('settings.path.title')}</Label>
<div class="grid grid-cols-2 justify-center gap-5 pt-5">
	<p class="text-sm text-gray-500">{$_('settings.path.ask_me_desc')}</p>
	<Toggle bind:checked={$askWheretoSave}>{$_('settings.path.ask_me')}</Toggle>
</div>

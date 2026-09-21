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
	import CpuSettings from '$lib/settings/CPUSettings.svelte';
	import ImageSettings from '$lib/settings/ImageSettings.svelte';
	import LangSelect from '$lib/settings/LangSelect.svelte';
	import PathSettings from '$lib/settings/PathSettings.svelte';
</script>

<main class="flex h-full min-h-0 w-full flex-col gap-4 overflow-y-auto pb-1">
	<div class="grid shrink-0 grid-cols-2 gap-4">
		<div class="rounded-xl border border-line bg-surface p-5">
			<LangSelect />
		</div>
		<div class="rounded-xl border border-line bg-surface p-5">
			<PathSettings />
		</div>
	</div>
	<div class="shrink-0 rounded-xl border border-line bg-surface p-5">
		<CpuSettings />
	</div>
	<div class="shrink-0 rounded-xl border border-line bg-surface p-5">
		<ImageSettings />
	</div>
</main>

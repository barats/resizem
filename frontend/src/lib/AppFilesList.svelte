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
	import { Listgroup, ListgroupItem } from 'flowbite-svelte';
	import { filesList, resultList } from '$lib/app_stores';
	import { FileImageOutline, CheckCircleOutline, CloseCircleOutline } from 'flowbite-svelte-icons';
</script>

<Listgroup class="w-full border-0">
	{#if $filesList.length > 0}
		{#each $filesList as item}
			<ListgroupItem class="flex gap-2 text-sm">
				<FileImageOutline color="w-5 h-5" />{item}
			</ListgroupItem>
		{/each}
	{/if}

	{#if $resultList.length > 0}
		{#each $resultList as item}
			<ListgroupItem class="flex gap-2 text-sm">
				{#if item.status === 1}
					<CheckCircleOutline class="h-5 w-5" />{item.name}
				{:else}
					<CloseCircleOutline class="h-5 w-5" />{item.name} - {item.message}
				{/if}
			</ListgroupItem>
		{/each}
	{/if}
</Listgroup>

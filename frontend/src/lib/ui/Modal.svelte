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
	let { open = $bindable(false), title = '', class: klass = '', children, footer } = $props();

	let dialogEl = $state(null);

	$effect(() => {
		if (!dialogEl) return;
		if (open && !dialogEl.open) {
			dialogEl.showModal();
		} else if (!open && dialogEl.open) {
			dialogEl.close();
		}
	});

	function handleClose() {
		open = false;
	}

	// Clicks on the <dialog> element itself (outside the panel) close it.
	function handleClick(event) {
		if (event.target === dialogEl) {
			open = false;
		}
	}
</script>

<dialog
	bind:this={dialogEl}
	onclose={handleClose}
	onclick={handleClick}
	aria-label={title}
	class="m-auto w-[calc(100vw-2rem)] max-w-lg rounded-xl bg-surface p-0 text-ink shadow-xl backdrop:bg-zinc-950/40 {klass}"
>
	<div class="flex max-h-[80vh] flex-col p-5">
		<h2 class="mb-3 shrink-0 text-base font-semibold">{title}</h2>
		<div
			class="min-h-0 flex-1 overflow-y-auto pr-1 text-sm text-ink-secondary [&_p]:mb-3 [&_strong]:text-ink"
		>
			{@render children?.()}
		</div>
		{#if footer}
			<div class="mt-4 flex shrink-0 justify-end">{@render footer()}</div>
		{/if}
	</div>
</dialog>

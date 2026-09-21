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
	/**
	 * @type {{
	 *   variant?: 'primary' | 'secondary' | 'ghost' | 'danger',
	 *   size?: 'sm' | 'md' | 'icon',
	 *   type?: 'button' | 'submit' | 'reset',
	 *   class?: string,
	 *   children?: import('svelte').Snippet
	 * } & Record<string, unknown>}
	 */
	let {
		variant = 'secondary',
		size = 'md',
		type = 'button',
		class: klass = '',
		children,
		...rest
	} = $props();

	const variants = {
		primary:
			'bg-primary-600 text-white hover:bg-primary-700 active:scale-[0.98] focus-visible:outline-primary-600',
		secondary:
			'bg-surface text-ink border border-line hover:bg-surface-muted active:scale-[0.98] focus-visible:outline-primary-600',
		ghost:
			'text-ink-secondary hover:bg-surface-muted hover:text-ink active:scale-[0.98] focus-visible:outline-primary-600',
		danger:
			'bg-red-600 text-white hover:bg-red-700 active:scale-[0.98] focus-visible:outline-red-600'
	};

	const sizes = {
		sm: 'h-8 gap-1.5 px-3 text-xs',
		md: 'h-9 gap-2 px-3.5 text-sm',
		icon: 'h-9 w-9 justify-center p-0'
	};
</script>

<button
	{type}
	class="inline-flex shrink-0 items-center justify-center rounded-lg font-medium whitespace-nowrap transition-[color,background-color,border-color,transform] duration-150 select-none focus-visible:outline-2 focus-visible:outline-offset-2 disabled:pointer-events-none disabled:opacity-60 {variants[
		variant
	] ?? variants.secondary} {sizes[size] ?? sizes.md} {klass}"
	{...rest}
>
	{@render children?.()}
</button>

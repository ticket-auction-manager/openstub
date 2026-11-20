<script lang="ts">
	import './layout.css';
	import favicon from '#lib/assets/sq.svg';
	import banner from '#lib/assets/banner.svg';
	import type { LayoutProps } from './$types';

	let pageTitle: string = "OpenStub";

	let { children, data }: LayoutProps = $props();
</script>

<svelte:head>
  <title>{pageTitle} - {data.deviceName}</title>
	<link rel="icon" href={favicon} />
</svelte:head>
{#await data then}
<div class="p-4 flex flex-row gap-4 bg-gray-300 dark:bg-gray-800">
  <img src={banner} alt="OpenStub logo" class="block h-8">
  <h1 class="text-xl font-bold">{pageTitle}</h1>
  <div>{data.deviceName}</div>
</div>
{@render children()}
{/await}

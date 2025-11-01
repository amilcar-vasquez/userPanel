<script lang="ts">
	import { onMount } from 'svelte';
	import { auth } from '$lib/stores/auth';
	import Toast from '$lib/components/Toast.svelte';
	import favicon from '$lib/assets/favicon.svg';
	import '../app.css';

	let { children } = $props();

	// Initialize auth state from localStorage on app load
	onMount(() => {
		// Refresh profile if token exists
		auth.subscribe((state) => {
			if (state.token && !state.user) {
				auth.fetchProfile();
			}
		});
	});
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<Toast />
{@render children()}

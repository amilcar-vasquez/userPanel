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

<header class="topbar">
	<div class="brand">Project Showcase</div>
	<nav class="nav">
		<a href="/">Home</a>
		<a href="/dashboard">Dashboard</a>
		<a href="/profiles">Profiles</a>
		<a href="/projects">Projects</a>
		<a href="/settings">Settings</a>
		<a href="/login">Login</a>
	</nav>
</header>

<main>
	<Toast />
	{@render children()}
</main>

<style>
	.topbar {
		display: flex;
		align-items: center;
		justify-content: space-between;
		padding: 12px 20px;
		border-bottom: 1px solid var(--md-sys-color-outline-variant);
		background: var(--md-sys-color-surface);
	}

	.brand {
		font-weight: 600;
		font-size: 18px;
		color: var(--md-sys-color-primary);
	}

	.nav a {
		margin-left: 12px;
		color: var(--md-sys-color-on-surface-variant);
		text-decoration: none;
	}

	.nav a:hover {
		color: var(--md-sys-color-on-surface);
		text-decoration: underline;
	}

	main {
		min-height: calc(100vh - 64px);
	}
</style>

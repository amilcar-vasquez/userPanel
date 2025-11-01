<script lang="ts">
	import { isAuthenticated } from '$lib/stores/auth';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import '@material/web/button/filled-button.js';
	import '@material/web/button/outlined-button.js';
	import '@material/web/icon/icon.js';

	let authenticated = false;
	isAuthenticated.subscribe((value) => (authenticated = value));

	onMount(() => {
		// Redirect to appropriate page
		if (authenticated) {
			goto('/profile');
		} else {
			goto('/login');
		}
	});
</script>

<svelte:head>
	<title>Auth Service</title>
</svelte:head>

<div class="container">
	<div class="hero">
		<md-icon class="hero-icon">lock</md-icon>
		<h1>Auth Service</h1>
		<p class="subtitle">Secure authentication for your applications</p>

		<div class="actions">
			<md-filled-button href="/login">
				<md-icon slot="icon">login</md-icon>
				Sign In
			</md-filled-button>
			<md-outlined-button href="/register">
				<md-icon slot="icon">person_add</md-icon>
				Create Account
			</md-outlined-button>
		</div>
	</div>
</div>

<style>
	.container {
		display: flex;
		justify-content: center;
		align-items: center;
		min-height: 100vh;
		padding: 24px;
		background: var(--md-sys-color-surface-container-low);
	}

	.hero {
		text-align: center;
		max-width: 600px;
	}

	.hero-icon {
		font-size: 96px;
		color: var(--md-sys-color-primary);
		margin-bottom: 24px;
	}

	h1 {
		font-size: 48px;
		font-weight: 500;
		color: var(--md-sys-color-on-surface);
		margin: 0 0 16px 0;
	}

	.subtitle {
		font-size: 20px;
		color: var(--md-sys-color-on-surface-variant);
		margin: 0 0 48px 0;
	}

	.actions {
		display: flex;
		gap: 16px;
		justify-content: center;
		flex-wrap: wrap;
	}

	@media (max-width: 600px) {
		h1 {
			font-size: 36px;
		}

		.subtitle {
			font-size: 18px;
		}

		.hero-icon {
			font-size: 72px;
		}
	}
</style>

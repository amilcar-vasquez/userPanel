<script lang="ts">
	import { auth, isAuthenticated } from '$lib/stores/auth';
	import { toast } from '$lib/stores/toast';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import '@material/web/textfield/outlined-text-field.js';
	import '@material/web/button/filled-button.js';
	import '@material/web/button/text-button.js';
	import '@material/web/progress/circular-progress.js';

	let email = '';
	let password = '';
	let loading = false;

	// Redirect if already authenticated
	onMount(() => {
		isAuthenticated.subscribe((authenticated) => {
			if (authenticated) {
				goto('/profile');
			}
		});
	});

	async function handleLogin() {
		if (!email || !password) {
			toast.error('Please fill in all fields');
			return;
		}

		loading = true;
		try {
			await auth.login(email, password);
			toast.success('Login successful!');
		} catch (error: any) {
			toast.error(error.message || 'Login failed');
		} finally {
			loading = false;
		}
	}

	function handleKeyPress(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			handleLogin();
		}
	}
</script>

<svelte:head>
	<title>Login - Auth Service</title>
</svelte:head>

<div class="container">
	<div class="card">
		<div class="header">
			<h1>Welcome Back</h1>
			<p class="subtitle">Sign in to your account</p>
		</div>

		<form on:submit|preventDefault={handleLogin}>
		<div class="form-field">
			<md-outlined-text-field
				label="Email"
				type="email"
				value={email}
				on:input={(e: any) => (email = e.target.value)}
				on:keypress={handleKeyPress}
				required
				style="width: 100%;"
			/>
		</div>		<div class="form-field">
			<md-outlined-text-field
				label="Password"
				type="password"
				value={password}
				on:input={(e: any) => (password = e.target.value)}
				on:keypress={handleKeyPress}
				required
				style="width: 100%;"
			/>
		</div>			<div class="actions">
				<md-filled-button type="submit" disabled={loading} style="width: 100%;">
					{#if loading}
						<md-circular-progress indeterminate slot="icon" style="--md-circular-progress-size: 20px;" />
						Signing in...
					{:else}
						Sign In
					{/if}
				</md-filled-button>
			</div>
		</form>

		<div class="footer">
			<span class="footer-text">Don't have an account?</span>
			<md-text-button href="/register">Create Account</md-text-button>
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

	.card {
		background: var(--md-sys-color-surface);
		border-radius: 28px;
		padding: 48px;
		max-width: 450px;
		width: 100%;
		box-shadow: var(--md-sys-elevation-1);
	}

	.header {
		text-align: center;
		margin-bottom: 32px;
	}

	h1 {
		font-size: 32px;
		font-weight: 500;
		color: var(--md-sys-color-on-surface);
		margin: 0 0 8px 0;
	}

	.subtitle {
		font-size: 16px;
		color: var(--md-sys-color-on-surface-variant);
		margin: 0;
	}

	.form-field {
		margin-bottom: 24px;
	}

	.actions {
		margin-top: 32px;
	}

	.footer {
		margin-top: 24px;
		text-align: center;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 8px;
	}

	.footer-text {
		font-size: 14px;
		color: var(--md-sys-color-on-surface-variant);
	}

	@media (max-width: 600px) {
		.card {
			padding: 32px 24px;
		}

		h1 {
			font-size: 28px;
		}
	}
</style>

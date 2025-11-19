<script lang="ts">
	import { auth, isAuthenticated } from '$lib/stores/auth';
	import { toast } from '$lib/stores/toast';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import '@material/web/textfield/outlined-text-field.js';
	import '@material/web/button/filled-button.js';
	import '@material/web/button/text-button.js';
	import '@material/web/progress/circular-progress.js';

	let name = '';
	let email = '';
	let password = '';
	let confirmPassword = '';
	let loading = false;

	// Redirect if already authenticated
	onMount(() => {
		isAuthenticated.subscribe((authenticated) => {
			if (authenticated) {
				goto('/profile');
			}
		});
	});

	async function handleRegister() {
		// Validation
		if (!name || !email || !password || !confirmPassword) {
			toast.error('Please fill in all fields');
			return;
		}

		if (password.length < 6) {
			toast.error('Password must be at least 6 characters');
			return;
		}

		if (password !== confirmPassword) {
			toast.error('Passwords do not match');
			return;
		}

		loading = true;
		try {
			await auth.register(name, email, password);
			toast.success('Account created successfully!');
		} catch (error: any) {
			toast.error(error.message || 'Registration failed');
		} finally {
			loading = false;
		}
	}

	function handleKeyPress(event: KeyboardEvent) {
		if (event.key === 'Enter') {
			handleRegister();
		}
	}
</script>

<svelte:head>
	<title>Register - Auth Service</title>
</svelte:head>

<div class="container">
	<div class="card">
		<div class="header">
			<h1>Create Account</h1>
			<p class="subtitle">Join us today</p>
		</div>

		<form on:submit|preventDefault={handleRegister}>
		<div class="form-field">
			<md-outlined-text-field
				label="Full Name"
				type="text"
				value={name}
				on:input={(e: any) => (name = e.target.value)}
				on:keypress={handleKeyPress}
				required
				style="width: 100%;"
			/>
		</div>		<div class="form-field">
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
				supporting-text="Minimum 6 characters"
				style="width: 100%;"
			/>
		</div>		<div class="form-field">
			<md-outlined-text-field
				label="Confirm Password"
				type="password"
				value={confirmPassword}
				on:input={(e: any) => (confirmPassword = e.target.value)}
				on:keypress={handleKeyPress}
				required
				style="width: 100%;"
			/>
		</div>			<div class="actions">
				<md-filled-button type="submit" disabled={loading} style="width: 100%;">
					{#if loading}
						<md-circular-progress indeterminate slot="icon" style="--md-circular-progress-size: 20px;" />
						Creating account...
					{:else}
						Create Account
					{/if}
				</md-filled-button>
			</div>
		</form>

		<div class="footer">
			<span class="footer-text">Already have an account?</span>
			<md-text-button href="/login">Sign In</md-text-button>
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
		margin-bottom: 20px;
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

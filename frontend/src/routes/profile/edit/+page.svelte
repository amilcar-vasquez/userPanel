<script lang="ts">
	import { auth, isAuthenticated } from '$lib/stores/auth';
	import { toast } from '$lib/stores/toast';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import '@material/web/textfield/outlined-text-field.js';
	import '@material/web/button/filled-button.js';
	import '@material/web/button/text-button.js';
	import '@material/web/progress/circular-progress.js';
	import '@material/web/icon/icon.js';
	import type { MdOutlinedTextField } from '@material/web/textfield/outlined-text-field.js';

	let nameField: MdOutlinedTextField;
	let authState: any;
	let loading = false;

	auth.subscribe((value) => (authState = value));

	onMount(() => {
		const unsubscribe = isAuthenticated.subscribe((authenticated) => {
			if (!authenticated) {
				goto('/login');
			}
		});

		return unsubscribe;
	});

	async function handleSave() {
		const name = nameField?.value || '';

		if (!name.trim()) {
			toast.error('Name cannot be empty');
			return;
		}

		loading = true;
		try {
			await auth.updateProfile({ name });
			toast.success('Profile updated successfully');
			goto('/profile');
		} catch (error: any) {
			toast.error(error.message || 'Failed to update profile');
		} finally {
			loading = false;
		}
	}

	function handleCancel() {
		goto('/profile');
	}
</script>

<svelte:head>
	<title>Edit Profile - Auth Service</title>
</svelte:head>

<div class="container">
	<div class="card">
		<div class="header">
			<h1>Edit Name</h1>
			<p class="subtitle">Update your full name</p>
		</div>

		{#if authState.user}
			<form on:submit|preventDefault={handleSave}>
				<div class="form-field">
					<md-outlined-text-field
						bind:this={nameField}
						label="Full Name"
						type="text"
						value={authState.user.name}
						required
						style="width: 100%;"
					>
						<md-icon slot="leading-icon">person</md-icon>
					</md-outlined-text-field>
				</div>

				<div class="form-field">
					<md-outlined-text-field
						label="Email"
						type="email"
						value={authState.user.email}
						disabled
						supporting-text="Email cannot be changed"
						style="width: 100%;"
					>
						<md-icon slot="leading-icon">email</md-icon>
					</md-outlined-text-field>
				</div>

				<div class="actions">
					<md-text-button type="button" on:click={handleCancel} style="flex: 1;">
						Cancel
					</md-text-button>
					<md-filled-button type="submit" disabled={loading} style="flex: 1;">
						{#if loading}
							<md-circular-progress
								indeterminate
								slot="icon"
								style="--md-circular-progress-size: 20px;"
							></md-circular-progress>
							Saving...
						{:else}
							<md-icon slot="icon">save</md-icon>
							Save Changes
						{/if}
					</md-filled-button>
				</div>
			</form>
		{/if}
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
		max-width: 500px;
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
		display: flex;
		gap: 12px;
	}

	@media (max-width: 600px) {
		.card {
			padding: 32px 24px;
		}

		h1 {
			font-size: 28px;
		}

		.actions {
			flex-direction: column-reverse;
		}
	}
</style>

<script lang="ts">
	import { auth, isAuthenticated } from '$lib/stores/auth';
	import { toast } from '$lib/stores/toast';
	import { goto } from '$app/navigation';
	import { onMount } from 'svelte';
	import '@material/web/button/filled-button.js';
	import '@material/web/button/outlined-button.js';
	import '@material/web/button/text-button.js';
	import '@material/web/icon/icon.js';
	import '@material/web/progress/circular-progress.js';
	import '@material/web/divider/divider.js';
	import '@material/web/dialog/dialog.js';
	import type { MdDialog } from '@material/web/dialog/dialog.js';

	let deleteDialog: MdDialog;
	let authState: any;
	let loading = true;

	auth.subscribe((value) => (authState = value));

	onMount(() => {
		// Redirect if not authenticated
		const unsubscribe = isAuthenticated.subscribe((authenticated) => {
			if (!authenticated) {
				goto('/login');
			}
		});

		// Fetch fresh profile data
		if (authState.token) {
			auth.fetchProfile().finally(() => {
				loading = false;
			});
		} else {
			loading = false;
		}

		return unsubscribe;
	});

	function handleLogout() {
		auth.logout();
		toast.info('You have been logged out');
	}

	function handleEdit() {
		goto('/profile/edit');
	}

	function confirmDelete() {
		deleteDialog.show();
	}

	async function handleDelete() {
		try {
			await auth.deleteAccount();
			toast.success('Account deleted successfully');
		} catch (error: any) {
			toast.error(error.message || 'Failed to delete account');
		}
	}

	function formatDate(dateString: string): string {
		const date = new Date(dateString);
		return date.toLocaleDateString('en-US', {
			year: 'numeric',
			month: 'long',
			day: 'numeric'
		});
	}
</script>

<svelte:head>
	<title>Profile - Auth Service</title>
</svelte:head>

<div class="container">
	<div class="profile-card">
		{#if loading}
			<div class="loading">
				<md-circular-progress indeterminate />
			</div>
		{:else if authState.user}
			<div class="profile-header">
				<h1 class="user-name">Account Information</h1>
				<p class="user-email">{authState.user.email}</p>
			</div>

			<md-divider></md-divider>

			<div class="profile-info">
				<div class="info-row">
					<span class="info-label">Account ID</span>
					<span class="info-value">{authState.user.id}</span>
				</div>

				<div class="info-row">
					<span class="info-label">Member Since</span>
					<span class="info-value">{formatDate(authState.user.created_at)}</span>
				</div>

				<div class="info-row">
					<span class="info-label">Last Updated</span>
					<span class="info-value">{formatDate(authState.user.updated_at)}</span>
				</div>
			</div>

			<md-divider></md-divider>

			<div class="profile-actions">
				<md-filled-button href="/profile/github" style="flex: 1;">
					<md-icon slot="icon">arrow_back</md-icon>
					Back to GitHub Profile
				</md-filled-button>
				<md-outlined-button on:click={handleEdit} style="flex: 1;">
					<md-icon slot="icon">edit</md-icon>
					Edit Name
				</md-outlined-button>
			</div>

			<div class="danger-zone">
				<md-text-button on:click={confirmDelete} style="color: var(--md-sys-color-error);">
					<md-icon slot="icon">delete</md-icon>
					Delete Account
				</md-text-button>
			</div>
		{:else}
			<div class="error">
				<p>Failed to load profile</p>
				<md-text-button on:click={() => goto('/login')}>Back to Login</md-text-button>
			</div>
		{/if}
	</div>
</div>

<!-- Delete Confirmation Dialog -->
<md-dialog bind:this={deleteDialog}>
	<div slot="headline">Delete Account?</div>
	<div slot="content">
		This action cannot be undone. Your account and all associated data will be permanently deleted.
	</div>
	<div slot="actions">
		<md-text-button on:click={() => deleteDialog.close()}>Cancel</md-text-button>
		<md-text-button
			on:click={() => {
				deleteDialog.close();
				handleDelete();
			}}
			style="color: var(--md-sys-color-error);"
		>
			Delete
		</md-text-button>
	</div>
</md-dialog>

<style>
	.container {
		display: flex;
		justify-content: center;
		align-items: center;
		min-height: 100vh;
		padding: 24px;
		background: var(--md-sys-color-surface-container-low);
	}

	.profile-card {
		background: var(--md-sys-color-surface);
		border-radius: 28px;
		padding: 48px;
		max-width: 600px;
		width: 100%;
		box-shadow: var(--md-sys-elevation-2);
	}

	.loading,
	.error {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 48px;
		text-align: center;
		color: var(--md-sys-color-on-surface-variant);
	}

	.profile-header {
		text-align: center;
		margin-bottom: 32px;
	}

	.user-name {
		font-size: 32px;
		font-weight: 500;
		color: var(--md-sys-color-on-surface);
		margin: 0 0 8px 0;
	}

	.user-email {
		font-size: 16px;
		color: var(--md-sys-color-on-surface-variant);
		margin: 0;
	}

	.profile-info {
		padding: 24px 0;
	}

	.info-row {
		display: flex;
		justify-content: space-between;
		align-items: center;
		padding: 12px 0;
	}

	.info-label {
		font-size: 14px;
		color: var(--md-sys-color-on-surface-variant);
		font-weight: 500;
	}

	.info-value {
		font-size: 14px;
		color: var(--md-sys-color-on-surface);
	}

	.profile-actions {
		display: flex;
		gap: 12px;
		margin-top: 24px;
		flex-wrap: wrap;
	}

	.danger-zone {
		margin-top: 32px;
		text-align: center;
		padding-top: 24px;
		border-top: 1px solid var(--md-sys-color-outline-variant);
	}

	@media (max-width: 600px) {
		.profile-card {
			padding: 32px 24px;
		}

		.user-name {
			font-size: 28px;
		}

		.profile-actions {
			flex-direction: column;
		}
	}
</style>

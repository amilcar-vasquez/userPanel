<script lang="ts">
	import { onMount } from 'svelte';
	import { auth, isAuthenticated } from '$lib/stores/auth';
	import { toast } from '$lib/stores/toast';
	import { goto } from '$app/navigation';
	import * as github from '$lib/github';
	import type { GitHubProfileStats, RankInfo, GitHubProfileResponse } from '$lib/github';
	import '@material/web/button/filled-button.js';
	import '@material/web/button/outlined-button.js';
	import '@material/web/button/text-button.js';
	import '@material/web/icon/icon.js';
	import '@material/web/progress/circular-progress.js';
	import '@material/web/divider/divider.js';
	import '@material/web/dialog/dialog.js';
	import '@material/web/textfield/outlined-text-field.js';
	import type { MdDialog } from '@material/web/dialog/dialog.js';
	import type { MdOutlinedTextField } from '@material/web/textfield/outlined-text-field.js';

	let githubStats: GitHubProfileStats | null = null;
	let rankInfo: RankInfo | null = null;
	let loading = true;
	let error: string | null = null;
	let authState: any;
	let settingsDialog: MdDialog;
	let deleteDialog: MdDialog;
	let githubUsernameField: MdOutlinedTextField;
	let githubTokenField: MdOutlinedTextField;
	let savingCredentials = false;

	auth.subscribe((value) => (authState = value));

	onMount(() => {
		// Redirect if not authenticated
		const unsubscribe = isAuthenticated.subscribe((authenticated) => {
			if (!authenticated) {
				goto('/login');
			}
		});

		loadGithubProfile();

		return unsubscribe;
	});

	async function loadGithubProfile() {
		loading = true;
		error = null;
		try {
			const response = await github.fetchGithubProfile();
			githubStats = response.profile;
			rankInfo = response.rank;
		} catch (err: any) {
			error = err.message || 'Failed to load GitHub profile';
			console.error('GitHub profile error:', err);
		} finally {
			loading = false;
		}
	}

	async function handleSaveCredentials() {
		const username = githubUsernameField.value;
		const token = githubTokenField.value;

		if (!username || !token) {
			toast.error('Please provide both GitHub username and token');
			return;
		}

		savingCredentials = true;
		try {
			await github.updateGithubCredentials(username, token);
			toast.success('GitHub credentials saved! Refreshing profile...');
			settingsDialog.close();
			// Refresh the auth state to get updated user info
			await auth.fetchProfile();
			// Then load the GitHub profile
			await loadGithubProfile();
		} catch (err: any) {
			toast.error(err.message || 'Failed to save credentials');
		} finally {
			savingCredentials = false;
		}
	}

	function handleLogout() {
		auth.logout();
		toast.info('You have been logged out');
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

	function getContributionLevel(count: number): string {
		if (count === 0) return 'none';
		if (count < 3) return 'low';
		if (count < 6) return 'medium';
		if (count < 9) return 'high';
		return 'very-high';
	}
</script>

<svelte:head>
	<title>GitHub Profile - Auth Service</title>
</svelte:head>

<div class="page-container">
	<div class="header">
		<h1 class="page-title">GitHub Profile</h1>
		<div class="header-actions">
			<md-outlined-button on:click={() => settingsDialog.show()}>
				<md-icon slot="icon">settings</md-icon>
				Settings
			</md-outlined-button>
			<md-text-button on:click={handleLogout}>
				<md-icon slot="icon">logout</md-icon>
				Logout
			</md-text-button>
		</div>
	</div>

	{#if loading}
		<div class="loading-container">
			<md-circular-progress indeterminate />
			<p>Loading GitHub profile...</p>
		</div>
	{:else if error}
		<div class="error-container">
			<md-icon class="error-icon">error</md-icon>
			<h2>Failed to Load Profile</h2>
			<p>{error}</p>
			<md-filled-button on:click={() => settingsDialog.show()}>
				<md-icon slot="icon">settings</md-icon>
				Configure GitHub Credentials
			</md-filled-button>
		</div>
	{:else if githubStats}
		<!-- Profile Header -->
		<div class="profile-header">
			<img src={githubStats.avatar_url} alt={githubStats.name} class="avatar" />
			<div class="profile-info">
				<h2 class="name">{githubStats.name}</h2>
				<p class="username">@{githubStats.login}</p>
				{#if githubStats.bio}
					<p class="bio">{githubStats.bio}</p>
				{/if}
			</div>
		</div>

		<!-- Rank Card -->
		{#if rankInfo}
			<div class="rank-section">
				<div class="rank-card">
					<div class="rank-header">
						<div class="rank-badge rank-{rankInfo.rank.toLowerCase().replace('+', 'plus')}">
							<span class="rank-tier">{rankInfo.rank}</span>
						</div>
						<div class="rank-info">
							<h3>Developer Rank</h3>
							<p class="rank-score">Score: {rankInfo.score.toLocaleString()}</p>
						</div>
					</div>
					{#if rankInfo.next_rank}
						<div class="rank-progress">
							<div class="progress-header">
								<span>Progress to {rankInfo.next_rank}</span>
								<span>{rankInfo.progress_percent}%</span>
							</div>
							<div class="progress-bar">
								<div class="progress-fill" style="width: {rankInfo.progress_percent}%"></div>
							</div>
							<p class="next-threshold">
								{(rankInfo.next_rank_threshold! - rankInfo.score).toLocaleString()} points to next rank
							</p>
						</div>
					{:else}
						<p class="max-rank">🎉 Maximum rank achieved!</p>
					{/if}
				</div>
			</div>
		{/if}

		<!-- Stats Cards -->
		<div class="stats-grid">
			<div class="stat-card">
				<md-icon class="stat-icon">commit</md-icon>
				<div class="stat-value">{githubStats.total_commits.toLocaleString()}</div>
				<div class="stat-label">Commits</div>
			</div>
			<div class="stat-card">
				<md-icon class="stat-icon">pull_request</md-icon>
				<div class="stat-value">{githubStats.total_pull_requests.toLocaleString()}</div>
				<div class="stat-label">Pull Requests</div>
			</div>
			<div class="stat-card">
				<md-icon class="stat-icon">bug_report</md-icon>
				<div class="stat-value">{githubStats.total_issues.toLocaleString()}</div>
				<div class="stat-label">Issues</div>
			</div>
			<div class="stat-card">
				<md-icon class="stat-icon">rate_review</md-icon>
				<div class="stat-value">{githubStats.total_reviews.toLocaleString()}</div>
				<div class="stat-label">Reviews</div>
			</div>
			<div class="stat-card highlight">
				<md-icon class="stat-icon">star</md-icon>
				<div class="stat-value">{githubStats.total_stars_earned.toLocaleString()}</div>
				<div class="stat-label">Stars Earned</div>
			</div>
			<div class="stat-card">
				<md-icon class="stat-icon">group</md-icon>
				<div class="stat-value">{githubStats.followers.toLocaleString()}</div>
				<div class="stat-label">Followers</div>
			</div>
		</div>

		<!-- Contribution Calendar -->
		<div class="section">
			<h3 class="section-title">
				{githubStats.contribution_calendar.total_contributions.toLocaleString()} contributions in the
				last year
			</h3>
			<div class="contribution-calendar">
				{#each githubStats.contribution_calendar.weeks as week}
					<div class="week">
						{#each week.contribution_days as day}
							<div
								class="day contribution-{getContributionLevel(day.contribution_count)}"
								title="{day.contribution_count} contributions on {day.date}"
							></div>
						{/each}
					</div>
				{/each}
			</div>
		</div>

		<!-- Pinned Repositories -->
		{#if githubStats.pinned_repositories.length > 0}
			<div class="section">
				<h3 class="section-title">Pinned Repositories</h3>
				<div class="repos-grid">
					{#each githubStats.pinned_repositories as repo}
						<a href={repo.url} target="_blank" rel="noopener noreferrer" class="repo-card">
							<div class="repo-header">
								<md-icon class="repo-icon">folder</md-icon>
								<h4 class="repo-name">{repo.name}</h4>
							</div>
							{#if repo.description}
								<p class="repo-description">{repo.description}</p>
							{/if}
							<div class="repo-stats">
								{#if repo.primary_language.name}
									<div class="repo-language">
										<span
											class="language-dot"
											style="background-color: {repo.primary_language.color || '#ccc'}"
										></span>
										{repo.primary_language.name}
									</div>
								{/if}
								<div class="repo-stat">
									<md-icon>star</md-icon>
									{repo.stargazer_count}
								</div>
								<div class="repo-stat">
									<md-icon>fork_right</md-icon>
									{repo.fork_count}
								</div>
							</div>
						</a>
					{/each}
				</div>
			</div>
		{/if}

		<!-- Account Actions -->
		<div class="danger-zone">
			<md-text-button on:click={confirmDelete} style="color: var(--md-sys-color-error);">
				<md-icon slot="icon">delete</md-icon>
				Delete Account
			</md-text-button>
		</div>
	{/if}
</div>

<!-- Settings Dialog -->
<md-dialog bind:this={settingsDialog}>
	<div slot="headline">GitHub Settings</div>
	<div slot="content" class="dialog-content">
		<p class="dialog-description">
			Enter your GitHub username and a Personal Access Token with <code>read:user</code> and
			<code>repo</code> permissions.
		</p>
		<md-outlined-text-field
			bind:this={githubUsernameField}
			label="GitHub Username"
			value={authState.user?.github_username || ''}
			style="width: 100%; margin-bottom: 16px;"
		>
			<md-icon slot="leading-icon">person</md-icon>
		</md-outlined-text-field>
		<md-outlined-text-field
			bind:this={githubTokenField}
			label="GitHub Personal Access Token"
			type="password"
			style="width: 100%;"
		>
			<md-icon slot="leading-icon">key</md-icon>
		</md-outlined-text-field>
		<p class="dialog-hint">
			<a
				href="https://github.com/settings/tokens/new"
				target="_blank"
				rel="noopener noreferrer"
			>
				Generate a token here
			</a>
		</p>
	</div>
	<div slot="actions">
		<md-text-button on:click={() => settingsDialog.close()}>Cancel</md-text-button>
		<md-filled-button on:click={handleSaveCredentials} disabled={savingCredentials}>
			{savingCredentials ? 'Saving...' : 'Save'}
		</md-filled-button>
	</div>
</md-dialog>

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
	.page-container {
		max-width: 1200px;
		margin: 0 auto;
		padding: 24px;
		min-height: 100vh;
		background: var(--md-sys-color-surface-container-low);
	}

	.header {
		display: flex;
		justify-content: space-between;
		align-items: center;
		margin-bottom: 32px;
		flex-wrap: wrap;
		gap: 16px;
	}

	.page-title {
		font-size: 32px;
		font-weight: 500;
		color: var(--md-sys-color-on-surface);
		margin: 0;
	}

	.header-actions {
		display: flex;
		gap: 8px;
		flex-wrap: wrap;
	}

	.loading-container,
	.error-container {
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		padding: 64px 24px;
		text-align: center;
		background: var(--md-sys-color-surface);
		border-radius: 28px;
		box-shadow: var(--md-sys-elevation-1);
	}

	.error-icon {
		font-size: 64px;
		color: var(--md-sys-color-error);
		margin-bottom: 16px;
	}

	.profile-header {
		display: flex;
		gap: 24px;
		align-items: center;
		background: var(--md-sys-color-surface);
		padding: 32px;
		border-radius: 28px;
		box-shadow: var(--md-sys-elevation-1);
		margin-bottom: 24px;
	}

	.avatar {
		width: 120px;
		height: 120px;
		border-radius: 50%;
		border: 4px solid var(--md-sys-color-primary-container);
	}

	.profile-info {
		flex: 1;
	}

	.name {
		font-size: 28px;
		font-weight: 500;
		color: var(--md-sys-color-on-surface);
		margin: 0 0 8px 0;
	}

	.username {
		font-size: 18px;
		color: var(--md-sys-color-on-surface-variant);
		margin: 0 0 12px 0;
	}

	.bio {
		font-size: 16px;
		color: var(--md-sys-color-on-surface);
		margin: 0;
	}

	.rank-section {
		margin-bottom: 32px;
	}

	.rank-card {
		background: linear-gradient(135deg, var(--md-sys-color-primary-container) 0%, var(--md-sys-color-tertiary-container) 100%);
		padding: 32px;
		border-radius: 24px;
		box-shadow: var(--md-sys-elevation-2);
	}

	.rank-header {
		display: flex;
		align-items: center;
		gap: 24px;
		margin-bottom: 24px;
	}

	.rank-badge {
		width: 80px;
		height: 80px;
		border-radius: 50%;
		display: flex;
		align-items: center;
		justify-content: center;
		font-size: 32px;
		font-weight: 700;
		box-shadow: var(--md-sys-elevation-3);
	}

	.rank-badge.rank-splus {
		background: linear-gradient(135deg, #ffd700 0%, #ffa500 100%);
		color: #000;
	}

	.rank-badge.rank-s {
		background: linear-gradient(135deg, #ff69b4 0%, #ff1493 100%);
		color: #fff;
	}

	.rank-badge.rank-aplus {
		background: linear-gradient(135deg, #00ced1 0%, #1e90ff 100%);
		color: #fff;
	}

	.rank-badge.rank-a {
		background: linear-gradient(135deg, #32cd32 0%, #228b22 100%);
		color: #fff;
	}

	.rank-badge.rank-bplus {
		background: linear-gradient(135deg, #ffd700 0%, #daa520 100%);
		color: #000;
	}

	.rank-badge.rank-b {
		background: linear-gradient(135deg, #c0c0c0 0%, #808080 100%);
		color: #000;
	}

	.rank-badge.rank-c {
		background: linear-gradient(135deg, #cd7f32 0%, #8b4513 100%);
		color: #fff;
	}

	.rank-tier {
		font-size: 32px;
		font-weight: 700;
	}

	.rank-info h3 {
		margin: 0 0 8px 0;
		font-size: 24px;
		font-weight: 500;
		color: var(--md-sys-color-on-primary-container);
	}

	.rank-score {
		margin: 0;
		font-size: 16px;
		color: var(--md-sys-color-on-tertiary-container);
		font-weight: 500;
	}

	.rank-progress {
		margin-top: 16px;
	}

	.progress-header {
		display: flex;
		justify-content: space-between;
		margin-bottom: 8px;
		font-size: 14px;
		font-weight: 500;
		color: var(--md-sys-color-on-primary-container);
	}

	.progress-bar {
		width: 100%;
		height: 12px;
		background: rgba(0, 0, 0, 0.1);
		border-radius: 8px;
		overflow: hidden;
	}

	.progress-fill {
		height: 100%;
		background: linear-gradient(90deg, var(--md-sys-color-primary) 0%, var(--md-sys-color-tertiary) 100%);
		transition: width 0.5s ease-in-out;
		border-radius: 8px;
	}

	.next-threshold {
		margin: 8px 0 0 0;
		font-size: 13px;
		color: var(--md-sys-color-on-tertiary-container);
	}

	.max-rank {
		margin: 0;
		font-size: 18px;
		font-weight: 500;
		color: var(--md-sys-color-on-primary-container);
		text-align: center;
	}

	.stats-grid {
		display: grid;
		grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
		gap: 16px;
		margin-bottom: 24px;
	}

	.stat-card {
		background: var(--md-sys-color-surface);
		padding: 24px;
		border-radius: 16px;
		box-shadow: var(--md-sys-elevation-1);
		display: flex;
		flex-direction: column;
		align-items: center;
		text-align: center;
		transition: transform 0.2s, box-shadow 0.2s;
	}

	.stat-card:hover {
		transform: translateY(-2px);
		box-shadow: var(--md-sys-elevation-2);
	}

	.stat-card.highlight {
		background: var(--md-sys-color-primary-container);
	}

	.stat-icon {
		font-size: 32px;
		color: var(--md-sys-color-primary);
		margin-bottom: 12px;
	}

	.stat-value {
		font-size: 32px;
		font-weight: 500;
		color: var(--md-sys-color-on-surface);
		margin-bottom: 4px;
	}

	.stat-label {
		font-size: 14px;
		color: var(--md-sys-color-on-surface-variant);
	}

	.section {
		background: var(--md-sys-color-surface);
		padding: 32px;
		border-radius: 28px;
		box-shadow: var(--md-sys-elevation-1);
		margin-bottom: 24px;
	}

	.section-title {
		font-size: 20px;
		font-weight: 500;
		color: var(--md-sys-color-on-surface);
		margin: 0 0 24px 0;
	}

	.contribution-calendar {
		display: flex;
		gap: 3px;
		overflow-x: auto;
		padding: 8px 0;
	}

	.week {
		display: flex;
		flex-direction: column;
		gap: 3px;
	}

	.day {
		width: 12px;
		height: 12px;
		border-radius: 2px;
		background: var(--md-sys-color-surface-container-highest);
	}

	.day.contribution-none {
		background: var(--md-sys-color-surface-container);
	}

	.day.contribution-low {
		background: rgba(var(--md-sys-color-primary-rgb, 103, 80, 164), 0.3);
	}

	.day.contribution-medium {
		background: rgba(var(--md-sys-color-primary-rgb, 103, 80, 164), 0.6);
	}

	.day.contribution-high {
		background: rgba(var(--md-sys-color-primary-rgb, 103, 80, 164), 0.85);
	}

	.day.contribution-very-high {
		background: var(--md-sys-color-primary);
	}

	.repos-grid {
		display: grid;
		grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
		gap: 16px;
	}

	.repo-card {
		background: var(--md-sys-color-surface-container);
		padding: 20px;
		border-radius: 16px;
		border: 1px solid var(--md-sys-color-outline-variant);
		text-decoration: none;
		color: inherit;
		transition: transform 0.2s, box-shadow 0.2s;
		display: flex;
		flex-direction: column;
	}

	.repo-card:hover {
		transform: translateY(-2px);
		box-shadow: var(--md-sys-elevation-2);
		border-color: var(--md-sys-color-primary);
	}

	.repo-header {
		display: flex;
		align-items: center;
		gap: 8px;
		margin-bottom: 12px;
	}

	.repo-icon {
		color: var(--md-sys-color-on-surface-variant);
	}

	.repo-name {
		font-size: 16px;
		font-weight: 500;
		color: var(--md-sys-color-primary);
		margin: 0;
	}

	.repo-description {
		font-size: 14px;
		color: var(--md-sys-color-on-surface-variant);
		margin: 0 0 16px 0;
		flex: 1;
		display: -webkit-box;
		-webkit-line-clamp: 2;
		-webkit-box-orient: vertical;
		overflow: hidden;
	}

	.repo-stats {
		display: flex;
		gap: 16px;
		align-items: center;
		font-size: 13px;
		color: var(--md-sys-color-on-surface-variant);
		flex-wrap: wrap;
	}

	.repo-language {
		display: flex;
		align-items: center;
		gap: 6px;
	}

	.language-dot {
		width: 12px;
		height: 12px;
		border-radius: 50%;
	}

	.repo-stat {
		display: flex;
		align-items: center;
		gap: 4px;
	}

	.repo-stat md-icon {
		font-size: 16px;
	}

	.danger-zone {
		margin-top: 48px;
		text-align: center;
		padding: 32px;
		border-top: 1px solid var(--md-sys-color-outline-variant);
	}

	.dialog-content {
		display: flex;
		flex-direction: column;
		gap: 8px;
		min-width: 400px;
	}

	.dialog-description {
		margin: 0 0 16px 0;
		color: var(--md-sys-color-on-surface-variant);
	}

	.dialog-description code {
		background: var(--md-sys-color-surface-container);
		padding: 2px 6px;
		border-radius: 4px;
		font-family: 'Courier New', monospace;
		font-size: 13px;
	}

	.dialog-hint {
		margin: 8px 0 0 0;
		font-size: 13px;
		color: var(--md-sys-color-on-surface-variant);
	}

	.dialog-hint a {
		color: var(--md-sys-color-primary);
		text-decoration: none;
	}

	.dialog-hint a:hover {
		text-decoration: underline;
	}

	@media (max-width: 768px) {
		.page-container {
			padding: 16px;
		}

		.header {
			flex-direction: column;
			align-items: flex-start;
		}

		.profile-header {
			flex-direction: column;
			text-align: center;
		}

		.stats-grid {
			grid-template-columns: repeat(2, 1fr);
		}

		.repos-grid {
			grid-template-columns: 1fr;
		}

		.dialog-content {
			min-width: unset;
		}
	}
</style>

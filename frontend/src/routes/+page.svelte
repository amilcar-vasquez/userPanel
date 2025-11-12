<script lang="ts">
	import '@material/web/button/filled-button.js';
	import '@material/web/button/outlined-button.js';
	import '@material/web/icon/icon.js';
</script>

<svelte:head>
	<title>Project Showcase — Home</title>
</svelte:head>

<div class="container">
	<div class="hero">
		<md-icon class="hero-icon">insights</md-icon>
		<h1>GitHub Profile Showcase</h1>
		<p class="subtitle">A minimal UI to showcase two GitHub profiles, projects, and analytics.</p>

		<div class="actions">
			<md-filled-button href="/dashboard">
				<md-icon slot="icon">dashboard</md-icon>
				View Dashboard
			</md-filled-button>
			<md-outlined-button href="/profiles">
				<md-icon slot="icon">people</md-icon>
				Browse Profiles
			</md-outlined-button>
		</div>

		<div class="quick-links">
			<a href="/profiles/alice">Profile: Alice</a>
			<a href="/profiles/bob">Profile: Bob</a>
			<a href="/projects">All Projects</a>
		</div>
	</div>
</div>

<style>
	.container {
		display: flex;
		justify-content: center;
		align-items: center;
		min-height: 80vh;
		padding: 24px;
	}

	.hero {
		text-align: center;
		max-width: 760px;
		background: var(--md-sys-color-surface);
		padding: 36px;
		border-radius: 16px;
		box-shadow: var(--md-sys-elevation-1);
	}

	.hero-icon {
		font-size: 72px;
		color: var(--md-sys-color-primary);
		margin-bottom: 16px;
	}

	h1 { margin: 0 0 8px 0; }

	.subtitle { margin: 0 0 20px 0; color: var(--md-sys-color-on-surface-variant); }

	.actions { display:flex; gap:12px; justify-content:center; margin-bottom: 16px; }

	.quick-links { display:flex; gap:12px; justify-content:center; margin-top: 8px; }

	.quick-links a { color: var(--md-sys-color-primary); text-decoration: none; }
</style>

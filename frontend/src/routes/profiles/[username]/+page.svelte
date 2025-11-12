<script lang="ts">
    import { page } from '$app/stores';
    import GraphPlaceholder from '$lib/components/GraphPlaceholder.svelte';

    $: username = $page.params.username || 'unknown';
</script>

<svelte:head>
    <title>Profile — {username}</title>
</svelte:head>

<div class="container">
    <h1>Profile: {username}</h1>
    <p class="lead">This is a reusable profile page for <strong>{username}</strong>. Replace with live GitHub data.</p>

    <div class="grid">
        <div class="card">
            <GraphPlaceholder title="Contributions" />
        </div>
        <div class="card">
            <GraphPlaceholder title="Top Projects" />
        </div>
    </div>

    <div style="margin-top:16px">
        <a href="/projects">See all projects</a>
    </div>
</div>

<style>
    .container { padding: 24px; }
    .lead { color: var(--md-sys-color-on-surface-variant); }
    .grid { display:grid; grid-template-columns: 1fr 1fr; gap:12px; margin-top: 12px; }
    .card { background: var(--md-sys-color-surface); padding:8px; border-radius:8px; box-shadow:var(--md-sys-elevation-1); }
    @media (max-width:600px) { .grid { grid-template-columns: 1fr; } }
</style>

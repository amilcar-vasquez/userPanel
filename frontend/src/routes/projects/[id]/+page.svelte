<script lang="ts">
    import { page } from '$app/stores';
    import GraphPlaceholder from '$lib/components/GraphPlaceholder.svelte';

    $: id = $page.params.id || 'unknown';
</script>

<svelte:head>
    <title>Project — {id}</title>
</svelte:head>

<div class="container">
    <h1>Project: {id}</h1>
    <p class="lead">Generic project page. Replace with repository metadata and analytics when available.</p>

    <div class="grid">
        <div class="card"><GraphPlaceholder title="Activity" /></div>
        <div class="card"><GraphPlaceholder title="Issues / PRs" /></div>
    </div>

    <div style="margin-top:12px"><a href="/projects">Back to projects</a></div>
</div>

<style>
    .container { padding:24px }
    .lead { color: var(--md-sys-color-on-surface-variant) }
    .grid { display:grid; grid-template-columns: 1fr 1fr; gap:12px }
    .card { background: var(--md-sys-color-surface); padding:8px; border-radius:8px }
</style>

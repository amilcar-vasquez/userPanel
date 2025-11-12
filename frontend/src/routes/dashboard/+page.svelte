<script lang="ts">
    import GraphPlaceholder from '$lib/components/GraphPlaceholder.svelte';
    import '@material/web/button/filled-button.js';
    import '@material/web/button/outlined-button.js';
</script>

<svelte:head>
    <title>Dashboard — Analytics</title>
</svelte:head>

<div class="container">
    <h1>Dashboard</h1>
    <p class="lead">Analytics and visualizations for the showcased profiles and projects (placeholders).</p>

    <div class="grid">
        <div class="card">
            <GraphPlaceholder title="Contributions Over Time" height="160px" />
        </div>
        <div class="card">
            <GraphPlaceholder title="Top Repositories" height="160px" />
        </div>
        <div class="card full">
            <GraphPlaceholder title="Language Breakdown" height="220px" />
        </div>
    </div>
</div>

<style>
    .container { padding: 24px; }
    h1 { margin: 0 0 6px 0; }
    .lead { color: var(--md-sys-color-on-surface-variant); margin-bottom: 18px; }
    .grid { display:grid; grid-template-columns: repeat(auto-fit, minmax(240px, 1fr)); gap: 16px; }
    .card { background: var(--md-sys-color-surface); padding: 12px; border-radius: 10px; box-shadow: var(--md-sys-elevation-1); }
    .card.full { grid-column: 1 / -1; }
</style>

<script lang="ts">
    import GraphPlaceholder from '$lib/components/GraphPlaceholder.svelte';
</script>

<svelte:head>
    <title>Bob — Profile</title>
</svelte:head>

<div class="container">
    <h1>Bob</h1>
    <p class="lead">Sample profile for Bob — front-end showcase.</p>

    <div class="grid">
        <div class="card"><GraphPlaceholder title="Bob — Contributions" /></div>
        <div class="card"><GraphPlaceholder title="Bob — Projects" /></div>
    </div>

    <div style="margin-top:12px"><a href="/projects/project-bob-1">Bob's Project 1</a></div>
</div>

<style>
    .container { padding:24px }
    .lead { color: var(--md-sys-color-on-surface-variant) }
    .grid { display:grid; grid-template-columns: 1fr 1fr; gap:12px; }
    .card { background:var(--md-sys-color-surface); padding:8px; border-radius:8px }
</style>

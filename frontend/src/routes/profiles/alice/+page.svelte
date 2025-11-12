<script lang="ts">
    import GraphPlaceholder from '$lib/components/GraphPlaceholder.svelte';
</script>

<svelte:head>
    <title>Alice — Profile</title>
</svelte:head>

<div class="container">
    <h1>Alice</h1>
    <p class="lead">Sample profile for Alice — front-end showcase.</p>

    <div class="grid">
        <div class="card"><GraphPlaceholder title="Alice — Contributions" /></div>
        <div class="card"><GraphPlaceholder title="Alice — Projects" /></div>
    </div>

    <div style="margin-top:12px"><a href="/projects/project-alice-1">Alice's Project 1</a></div>
</div>

<style>
    .container { padding:24px }
    .lead { color: var(--md-sys-color-on-surface-variant) }
    .grid { display:grid; grid-template-columns: 1fr 1fr; gap:12px; }
    .card { background:var(--md-sys-color-surface); padding:8px; border-radius:8px }
</style>

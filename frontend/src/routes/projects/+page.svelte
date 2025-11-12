<script lang="ts">
    // Simple projects index with links to individual project pages
</script>

<svelte:head>
    <title>Projects</title>
</svelte:head>

<div class="container">
    <h1>Projects</h1>
    <p class="lead">List of sample projects for the showcased profiles.</p>

    <ul class="projects">
        <li><a href="/projects/project-alice-1">project-alice-1</a> — Alice</li>
        <li><a href="/projects/project-bob-1">project-bob-1</a> — Bob</li>
        <li><a href="/projects/shared-lib">shared-lib</a> — Shared</li>
    </ul>
</div>

<style>
    .container { padding:24px }
    .lead { color: var(--md-sys-color-on-surface-variant); margin-bottom:12px }
    .projects { list-style:none; padding:0; display:flex; flex-direction:column; gap:8px }
    a { color: var(--md-sys-color-primary) }
</style>

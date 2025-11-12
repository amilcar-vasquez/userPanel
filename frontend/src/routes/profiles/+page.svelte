<script lang="ts">
    import '@material/web/button/outlined-button.js';
</script>

<svelte:head>
    <title>Profiles — Showcase</title>
</svelte:head>

<div class="container">
    <h1>Profiles</h1>
    <p class="lead">Browse the sample profiles for the two showcased people.</p>

    <div class="list">
        <a class="profile-link" href="/profiles/alice">Alice</a>
        <a class="profile-link" href="/profiles/bob">Bob</a>
    </div>
</div>

<style>
    .container { padding: 24px; }
    .lead { color: var(--md-sys-color-on-surface-variant); margin-bottom: 12px; }
    .list { display:flex; gap:12px; }
    .profile-link { padding:10px 14px; background:var(--md-sys-color-surface); border-radius:8px; text-decoration:none; color:var(--md-sys-color-on-surface); box-shadow:var(--md-sys-elevation-1); }
</style>

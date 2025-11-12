<script lang="ts">
    import '@material/web/textfield/outlined-text-field.js';
    import '@material/web/button/filled-button.js';
</script>

<svelte:head>
    <title>Settings — Preferences</title>
</svelte:head>

<div class="container">
    <h1>Settings</h1>
    <p class="lead">User preferences and simple settings for the showcase app.</p>

    <form class="form">
        <label>Theme (placeholder)</label>
        <div style="margin:8px 0 16px 0; color:var(--md-sys-color-on-surface-variant);">Light / Dark toggle will go here.</div>

        <label>Notifications</label>
        <div style="margin:8px 0 16px 0; color:var(--md-sys-color-on-surface-variant);">Email / Push toggles will go here.</div>

        <md-filled-button>Save Preferences</md-filled-button>
    </form>
</div>

<style>
    .container { padding: 24px; }
    .lead { color: var(--md-sys-color-on-surface-variant); margin-bottom: 16px; }
    .form { max-width: 640px; }
</style>

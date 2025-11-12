<script lang="ts">
    export let title: string = 'Graph';
    export let height: string = '180px';
</script>

<div class="graph-root">
    <div class="graph-title">{title}</div>
    <div class="graph-box" style="height: {height};">Graph placeholder — data will render here.</div>
</div>

<style>
    .graph-root { padding: 12px; }
    .graph-title { font-weight: 600; margin-bottom: 8px; color: var(--md-sys-color-on-surface); }
    .graph-box {
        background: linear-gradient(90deg, rgba(0,0,0,0.02), rgba(0,0,0,0.01));
        border: 1px dashed var(--md-sys-color-outline-variant);
        border-radius: 8px;
        display:flex;
        align-items:center;
        justify-content:center;
        color: var(--md-sys-color-on-surface-variant);
    }
</style>

<script lang="ts">
	import { toast, type Toast } from '$lib/stores/toast';
	import { fly } from 'svelte/transition';
	import '@material/web/icon/icon.js';

	let toasts: Toast[] = [];
	toast.subscribe((value) => (toasts = value));

	function getIcon(type: Toast['type']): string {
		switch (type) {
			case 'success':
				return 'check_circle';
			case 'error':
				return 'error';
			default:
				return 'info';
		}
	}

	function getColor(type: Toast['type']): string {
		switch (type) {
			case 'success':
				return 'var(--md-sys-color-tertiary)';
			case 'error':
				return 'var(--md-sys-color-error)';
			default:
				return 'var(--md-sys-color-primary)';
		}
	}
</script>

<div class="toast-container">
	{#each toasts as t (t.id)}
		<div
			class="toast"
			style="--toast-color: {getColor(t.type)}"
			transition:fly={{ y: 50, duration: 300 }}
		>
			<md-icon class="toast-icon">{getIcon(t.type)}</md-icon>
			<span class="toast-message">{t.message}</span>
			<button class="toast-close" on:click={() => toast.dismiss(t.id)} aria-label="Close">
				<md-icon>close</md-icon>
			</button>
		</div>
	{/each}
</div>

<style>
	.toast-container {
		position: fixed;
		bottom: 24px;
		left: 50%;
		transform: translateX(-50%);
		z-index: 9999;
		display: flex;
		flex-direction: column;
		gap: 12px;
		pointer-events: none;
	}

	.toast {
		pointer-events: auto;
		display: flex;
		align-items: center;
		gap: 12px;
		padding: 12px 16px;
		background: var(--md-sys-color-surface-container-high);
		color: var(--md-sys-color-on-surface);
		border-radius: 8px;
		box-shadow: var(--md-sys-elevation-3);
		border-left: 4px solid var(--toast-color);
		min-width: 300px;
		max-width: 500px;
	}

	.toast-icon {
		color: var(--toast-color);
		font-size: 20px;
		flex-shrink: 0;
	}

	.toast-message {
		flex: 1;
		font-size: 14px;
		line-height: 20px;
	}

	.toast-close {
		background: none;
		border: none;
		cursor: pointer;
		padding: 4px;
		display: flex;
		align-items: center;
		justify-content: center;
		border-radius: 50%;
		color: var(--md-sys-color-on-surface-variant);
		transition: background 0.2s;
	}

	.toast-close:hover {
		background: var(--md-sys-color-surface-container-highest);
	}

	.toast-close md-icon {
		font-size: 18px;
	}
</style>

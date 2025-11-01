/**
 * Toast notification store for displaying snackbar messages
 */

import { writable } from 'svelte/store';

export interface Toast {
	id: number;
	message: string;
	type: 'success' | 'error' | 'info';
	duration?: number;
}

function createToastStore() {
	const { subscribe, update } = writable<Toast[]>([]);
	let idCounter = 0;

	return {
		subscribe,

		/**
		 * Show a toast notification
		 */
		show(message: string, type: Toast['type'] = 'info', duration = 4000) {
			const id = ++idCounter;
			const toast: Toast = { id, message, type, duration };

			update((toasts) => [...toasts, toast]);

			// Auto-dismiss after duration
			if (duration > 0) {
				setTimeout(() => {
					this.dismiss(id);
				}, duration);
			}

			return id;
		},

		/**
		 * Show success toast
		 */
		success(message: string, duration?: number) {
			return this.show(message, 'success', duration);
		},

		/**
		 * Show error toast
		 */
		error(message: string, duration?: number) {
			return this.show(message, 'error', duration);
		},

		/**
		 * Show info toast
		 */
		info(message: string, duration?: number) {
			return this.show(message, 'info', duration);
		},

		/**
		 * Dismiss a specific toast
		 */
		dismiss(id: number) {
			update((toasts) => toasts.filter((t) => t.id !== id));
		},

		/**
		 * Clear all toasts
		 */
		clear() {
			update(() => []);
		}
	};
}

export const toast = createToastStore();

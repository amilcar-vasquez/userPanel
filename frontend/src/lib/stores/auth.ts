/**
 * Authentication store for managing user state and JWT tokens
 * Persists to localStorage for session continuity
 */

import { writable, derived, get } from 'svelte/store';
import { goto } from '$app/navigation';
import { browser } from '$app/environment';
import * as api from '$lib/api';

export interface User {
	id: number;
	name: string;
	email: string;
	avatar?: string;
	github_username?: string;
	created_at: string;
	updated_at: string;
}

export interface AuthState {
	token: string | null;
	user: User | null;
	loading: boolean;
	error: string | null;
}

const initialState: AuthState = {
	token: null,
	user: null,
	loading: false,
	error: null
};

// Create the writable store
function createAuthStore() {
	const { subscribe, set, update } = writable<AuthState>(initialState);

	// Load state from localStorage on initialization
	if (browser) {
		const token = localStorage.getItem('auth_token');
		const userStr = localStorage.getItem('auth_user');
		if (token && userStr) {
			try {
				const user = JSON.parse(userStr);
				set({ token, user, loading: false, error: null });
			} catch (e) {
				// Invalid stored data, clear it
				localStorage.removeItem('auth_token');
				localStorage.removeItem('auth_user');
			}
		}
	}

	return {
		subscribe,

		/**
		 * Register a new user
		 */
		async register(name: string, email: string, password: string): Promise<void> {
			update((state) => ({ ...state, loading: true, error: null }));

			try {
				const response = await api.post<{ token: string; user: User }>('/register', {
					name,
					email,
					password
				});

				if (response.success && response.data) {
					const { token, user } = response.data;

					// Save to localStorage
					if (browser) {
						localStorage.setItem('auth_token', token);
						localStorage.setItem('auth_user', JSON.stringify(user));
					}

					// Update store
					set({ token, user, loading: false, error: null });

					// Redirect to profile
					goto('/profile');
				}
			} catch (error) {
				const errorMessage =
					error instanceof api.ApiError ? error.message : 'Registration failed';
				update((state) => ({ ...state, loading: false, error: errorMessage }));
				throw error;
			}
		},

		/**
		 * Login with email and password
		 */
		async login(email: string, password: string): Promise<void> {
			update((state) => ({ ...state, loading: true, error: null }));

			try {
				const response = await api.post<{ token: string; user: User }>('/login', {
					email,
					password
				});

				if (response.success && response.data) {
					const { token, user } = response.data;

					// Save to localStorage
					if (browser) {
						localStorage.setItem('auth_token', token);
						localStorage.setItem('auth_user', JSON.stringify(user));
					}

					// Update store
					set({ token, user, loading: false, error: null });

					// Redirect to profile
					goto('/profile');
				}
			} catch (error) {
				const errorMessage = error instanceof api.ApiError ? error.message : 'Login failed';
				update((state) => ({ ...state, loading: false, error: errorMessage }));
				throw error;
			}
		},

		/**
		 * Logout and clear session
		 */
		logout(): void {
			// Clear localStorage
			if (browser) {
				localStorage.removeItem('auth_token');
				localStorage.removeItem('auth_user');
			}

			// Reset store
			set(initialState);

			// Redirect to login
			goto('/login');
		},

		/**
		 * Fetch current user profile from API
		 */
		async fetchProfile(): Promise<void> {
			update((state) => ({ ...state, loading: true, error: null }));

			try {
				const response = await api.get<User>('/profile');

				if (response.success && response.data) {
					const user = response.data;

					// Update localStorage
					if (browser) {
						localStorage.setItem('auth_user', JSON.stringify(user));
					}

					// Update store
					update((state) => ({ ...state, user, loading: false }));
				}
			} catch (error) {
				const errorMessage =
					error instanceof api.ApiError ? error.message : 'Failed to fetch profile';

				// If unauthorized, clear auth and redirect to login
				if (error instanceof api.ApiError && error.status === 401) {
					if (browser) {
						localStorage.removeItem('auth_token');
						localStorage.removeItem('auth_user');
					}
					set(initialState);
					goto('/login');
				} else {
					update((state) => ({ ...state, loading: false, error: errorMessage }));
				}
			}
		},

		/**
		 * Update user profile
		 */
		async updateProfile(data: { name?: string; avatar?: string }): Promise<void> {
			update((state) => ({ ...state, loading: true, error: null }));

			try {
				const response = await api.put<User>('/profile', data);

				if (response.success && response.data) {
					const user = response.data;

					// Update localStorage
					if (browser) {
						localStorage.setItem('auth_user', JSON.stringify(user));
					}

					// Update store
					update((state) => ({ ...state, user, loading: false }));
				}
			} catch (error) {
				const errorMessage =
					error instanceof api.ApiError ? error.message : 'Failed to update profile';
				update((state) => ({ ...state, loading: false, error: errorMessage }));
				throw error;
			}
		},

		/**
		 * Delete user account
		 */
		async deleteAccount(): Promise<void> {
			update((state) => ({ ...state, loading: true, error: null }));

			try {
				await api.del('/profile');

				// Clear localStorage and reset store
				if (browser) {
					localStorage.removeItem('auth_token');
					localStorage.removeItem('auth_user');
				}
				set(initialState);

				// Redirect to login
				goto('/login');
			} catch (error) {
				const errorMessage =
					error instanceof api.ApiError ? error.message : 'Failed to delete account';
				update((state) => ({ ...state, loading: false, error: errorMessage }));
				throw error;
			}
		},

		/**
		 * Clear error message
		 */
		clearError(): void {
			update((state) => ({ ...state, error: null }));
		}
	};
}

export const auth = createAuthStore();

// Derived store for checking if user is authenticated
export const isAuthenticated = derived(auth, ($auth) => !!$auth.token && !!$auth.user);

// Derived store for getting user initials (for avatar)
export const userInitials = derived(auth, ($auth) => {
	if (!$auth.user) return '';
	const names = $auth.user.name.split(' ');
	if (names.length >= 2) {
		return `${names[0][0]}${names[1][0]}`.toUpperCase();
	}
	return $auth.user.name.substring(0, 2).toUpperCase();
});

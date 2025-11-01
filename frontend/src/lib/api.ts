/**
 * API client for communicating with the Go Auth Service
 * Base URL can be configured via PUBLIC_API_URL environment variable
 */

const API_BASE_URL = import.meta.env.PUBLIC_API_URL || 'http://localhost:8080/api';

export interface ApiResponse<T = any> {
	success: boolean;
	data?: T;
	message?: string;
}

export class ApiError extends Error {
	constructor(
		message: string,
		public status: number,
		public response?: any
	) {
		super(message);
		this.name = 'ApiError';
	}
}

/**
 * Get the stored JWT token from localStorage
 */
function getToken(): string | null {
	if (typeof window === 'undefined') return null;
	return localStorage.getItem('auth_token');
}

/**
 * Generic fetch wrapper with error handling
 */
async function fetchWithAuth(
	endpoint: string,
	options: RequestInit = {}
): Promise<ApiResponse> {
	const token = getToken();
	const url = `${API_BASE_URL}${endpoint}`;

	const headers: Record<string, string> = {
		'Content-Type': 'application/json'
	};

	// Add Authorization header if token exists
	if (token) {
		headers['Authorization'] = `Bearer ${token}`;
	}

	// Merge with custom headers from options
	if (options.headers) {
		Object.assign(headers, options.headers);
	}

	try {
		const response = await fetch(url, {
			...options,
			headers
		});

		const data = await response.json();

		if (!response.ok) {
			throw new ApiError(data.message || 'Request failed', response.status, data);
		}

		return data;
	} catch (error) {
		if (error instanceof ApiError) {
			throw error;
		}
		// Network or parsing error
		throw new ApiError(error instanceof Error ? error.message : 'Network error', 0);
	}
}

/**
 * GET request
 */
export async function get<T = any>(endpoint: string): Promise<ApiResponse<T>> {
	return fetchWithAuth(endpoint, {
		method: 'GET'
	});
}

/**
 * POST request
 */
export async function post<T = any>(endpoint: string, body?: any): Promise<ApiResponse<T>> {
	return fetchWithAuth(endpoint, {
		method: 'POST',
		body: body ? JSON.stringify(body) : undefined
	});
}

/**
 * PUT request
 */
export async function put<T = any>(endpoint: string, body?: any): Promise<ApiResponse<T>> {
	return fetchWithAuth(endpoint, {
		method: 'PUT',
		body: body ? JSON.stringify(body) : undefined
	});
}

/**
 * DELETE request
 */
export async function del<T = any>(endpoint: string): Promise<ApiResponse<T>> {
	return fetchWithAuth(endpoint, {
		method: 'DELETE'
	});
}

/**
 * Health check endpoint
 */
export async function checkHealth() {
	return get('/health');
}

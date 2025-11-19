import * as api from './api';

export interface GitHubProfileStats {
	login: string;
	name: string;
	avatar_url: string;
	bio: string;
	total_commits: number;
	total_pull_requests: number;
	total_issues: number;
	total_reviews: number;
	total_stars_earned: number;
	followers: number;
	contribution_calendar: ContributionCalendar;
	pinned_repositories: Repository[];
	total_public_repositories: number;
}

export interface ContributionCalendar {
	total_contributions: number;
	weeks: ContributionWeek[];
}

export interface ContributionWeek {
	contribution_days: ContributionDay[];
}

export interface ContributionDay {
	color: string;
	contribution_count: number;
	date: string;
}

export interface Repository {
	name: string;
	description: string;
	stargazer_count: number;
	fork_count: number;
	primary_language: Language;
	url: string;
}

export interface Language {
	name: string;
	color: string;
}

export interface RankInfo {
	rank: string;
	score: number;
	next_rank?: string;
	next_rank_threshold?: number;
	progress_percent: number;
}

export interface GitHubProfileResponse {
	profile: GitHubProfileStats;
	rank: RankInfo;
}

/**
 * Fetch GitHub profile stats and rank for the authenticated user
 */
export async function fetchGithubProfile(): Promise<GitHubProfileResponse> {
	const response = await api.get<GitHubProfileResponse>('/github/profile');
	return response.data!;
}

/**
 * Update GitHub credentials (username and token)
 */
export async function updateGithubCredentials(
	githubUsername: string,
	githubToken: string
): Promise<void> {
	await api.put('/github/credentials', {
		github_username: githubUsername,
		github_token: githubToken
	});
}

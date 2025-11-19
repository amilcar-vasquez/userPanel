# GitHub Profile Integration

This document describes the GitHub profile integration feature that allows users to display their GitHub statistics and contributions within the Auth Service.

## Overview

The GitHub integration uses the GitHub GraphQL API v4 to fetch comprehensive user profile statistics including:
- Contribution history and calendar
- Repository statistics
- Pinned repositories
- Stars earned across all repositories
- Commit, PR, and issue counts

## Architecture

### Backend

#### 1. Database Schema

Added fields to the `users` table:
- `github_username` (string): The user's GitHub username
- `github_token` (string): GitHub Personal Access Token (encrypted, never exposed in API responses)

#### 2. GitHub Client (`backend/internal/github/client.go`)

**Function: `FetchUserProfile(username, token string) (*UserProfileStats, error)`**

Uses the `shurcooL/githubv4` library to query GitHub's GraphQL API with the following data:

```graphql
query($username: String!, $from: DateTime!, $to: DateTime!) {
  user(login: $username) {
    login
    name
    avatarUrl
    bio
    contributionsCollection(from: $from, to: $to) {
      totalCommitContributions
      totalPullRequestContributions
      totalIssueContributions
      contributionCalendar {
        totalContributions
        weeks {
          contributionDays {
            color
            contributionCount
            date
          }
        }
      }
    }
    pinnedItems(first: 6, types: REPOSITORY) {
      nodes {
        ... on Repository {
          name
          description
          stargazerCount
          forkCount
          primaryLanguage {
            name
            color
          }
          url
        }
      }
    }
    repositories(first: 100, orderBy: {field: STARGAZERS, direction: DESC}, ownerAffiliations: OWNER, privacy: PUBLIC) {
      totalCount
      nodes {
        stargazerCount
      }
    }
  }
}
```

**Return Type: `UserProfileStats`**

```go
type UserProfileStats struct {
    Login                       string
    Name                        string
    AvatarURL                   string
    Bio                         string
    TotalCommits                int
    TotalPullRequests           int
    TotalIssues                 int
    TotalStarsEarned            int // Aggregated from all repos
    ContributionCalendar        ContributionCalendar
    PinnedRepositories          []Repository
    TotalPublicRepositories     int
}
```

#### 3. API Endpoints

**GET `/api/github/profile`** (Protected)
- Requires: JWT authentication
- Returns: GitHub profile statistics for the authenticated user
- Error cases:
  - 400: GitHub credentials not configured
  - 404: User not found
  - 500: GitHub API error

**PUT `/api/github/credentials`** (Protected)
- Requires: JWT authentication
- Body:
  ```json
  {
    "github_username": "username",
    "github_token": "ghp_xxx..."
  }
  ```
- Updates the user's GitHub credentials
- Returns: Success message with username

### Frontend

#### 1. GitHub API Module (`frontend/src/lib/github.ts`)

Provides TypeScript interfaces and API client functions:
- `fetchGithubProfile()`: Fetches stats from backend
- `updateGithubCredentials(username, token)`: Updates credentials

#### 2. GitHub Profile Page (`frontend/src/routes/profile/github/+page.svelte`)

A comprehensive GitHub-style profile page featuring:

**Profile Header**
- Avatar with primary color border
- Name and username
- Bio

**Statistics Cards**
- Total commits (this year)
- Total pull requests (this year)
- Total issues (this year)
- Total stars earned (highlighted card)

**Contribution Calendar**
- Visual heatmap of daily contributions
- Color-coded by contribution level (0-9+)
- Hover tooltips showing count and date

**Pinned Repositories**
- Grid of up to 6 pinned repos
- Repository name, description
- Primary language with color indicator
- Star and fork counts
- Clickable links to GitHub

**Settings Dialog**
- Form to enter GitHub username
- Secure password field for Personal Access Token
- Link to GitHub token generation page

## Setup Instructions

### 1. Generate a GitHub Personal Access Token

1. Go to https://github.com/settings/tokens/new
2. Give it a descriptive name (e.g., "Auth Service Integration")
3. Select the following scopes:
   - `read:user` - Read user profile data
   - `repo` - Access repository information
4. Click "Generate token"
5. **Copy the token immediately** (you won't be able to see it again)

### 2. Configure in the Application

1. Log in to your account
2. Go to Profile page
3. Click "GitHub Profile" button
4. Click "Settings" button in the top right
5. Enter your GitHub username
6. Paste your Personal Access Token
7. Click "Save"

### 3. View Your GitHub Profile

Once configured, the page will automatically load and display:
- Your contribution statistics for the current year
- A visual contribution calendar
- Your pinned repositories
- Total stars earned across all public repositories

## Security Considerations

1. **Token Storage**: GitHub tokens are stored in the database with the `json:"-"` tag, ensuring they're never exposed in API responses

2. **Authentication**: All GitHub endpoints require JWT authentication

3. **Scope Limitation**: The integration only requires `read:user` and `repo` scopes (read-only access)

4. **Client-Side**: Tokens are only transmitted during the initial setup and are not stored in browser localStorage or cookies

5. **API Rate Limits**: GitHub's GraphQL API has rate limits. The integration makes a single query per page load.

## Dependencies

### Backend
- `github.com/shurcooL/githubv4` v0.0.0-20240727222349-48295856cce7
- `github.com/shurcooL/graphql` v0.0.0-20230722043721-ed46e5a46466
- `golang.org/x/oauth2` v0.24.0

### Frontend
- Material Web Components (@material/web)
- SvelteKit for routing and SSR

## API Response Examples

### Successful Profile Fetch

```json
{
  "success": true,
  "data": {
    "login": "octocat",
    "name": "The Octocat",
    "avatar_url": "https://avatars.githubusercontent.com/u/583231",
    "bio": "GitHub mascot",
    "total_commits": 342,
    "total_pull_requests": 156,
    "total_issues": 89,
    "total_stars_earned": 45213,
    "contribution_calendar": {
      "total_contributions": 587,
      "weeks": [...]
    },
    "pinned_repositories": [
      {
        "name": "hello-world",
        "description": "My first repository",
        "stargazer_count": 1523,
        "fork_count": 234,
        "primary_language": {
          "name": "JavaScript",
          "color": "#f1e05a"
        },
        "url": "https://github.com/octocat/hello-world"
      }
    ],
    "total_public_repositories": 42
  }
}
```

### Error: Credentials Not Configured

```json
{
  "success": false,
  "message": "GitHub username or token not configured. Please update your profile first."
}
```

## Troubleshooting

### "Failed to fetch GitHub profile"

**Possible causes:**
1. Invalid GitHub token - Generate a new token
2. Token expired - Generate a new token
3. Incorrect username - Verify your GitHub username
4. Missing scopes - Ensure token has `read:user` and `repo` scopes
5. GitHub API rate limit - Wait and try again later

### Contribution calendar not showing

- The calendar shows contributions from the start of the current year
- If you just created your GitHub account this year, data may be limited

### Pinned repositories not showing

- You need to pin repositories on your GitHub profile
- Go to https://github.com/yourusername and click "Customize your pins"

## Future Enhancements

Potential improvements:
1. Cache GitHub data to reduce API calls
2. Webhook integration for real-time updates
3. Language statistics breakdown
4. Repository activity timeline
5. Contribution streak tracking
6. Compare with other users
7. Export statistics as image/PDF

## Testing

To test the integration locally:

1. Start the services:
   ```bash
   docker compose up -d
   ```

2. Register/login at http://localhost:3000

3. Navigate to Profile → GitHub Profile

4. Configure your credentials in Settings

5. The page should load your GitHub statistics

## Support

For issues or questions:
- Check GitHub token has correct scopes
- Verify username is correct
- Check browser console for error details
- Review backend logs: `docker compose logs api`

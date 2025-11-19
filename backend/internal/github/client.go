package github

import (
	"context"
	"time"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

// UserProfileStats represents aggregated GitHub profile statistics
type UserProfileStats struct {
	Login                   string               `json:"login"`
	Name                    string               `json:"name"`
	AvatarURL               string               `json:"avatar_url"`
	Bio                     string               `json:"bio"`
	TotalCommits            int                  `json:"total_commits"`
	TotalPullRequests       int                  `json:"total_pull_requests"`
	TotalIssues             int                  `json:"total_issues"`
	TotalStarsEarned        int                  `json:"total_stars_earned"`
	ContributionCalendar    ContributionCalendar `json:"contribution_calendar"`
	PinnedRepositories      []Repository         `json:"pinned_repositories"`
	TotalPublicRepositories int                  `json:"total_public_repositories"`
}

// ContributionCalendar represents the contribution calendar data
type ContributionCalendar struct {
	TotalContributions int                `json:"total_contributions"`
	Weeks              []ContributionWeek `json:"weeks"`
}

// ContributionWeek represents a week of contributions
type ContributionWeek struct {
	ContributionDays []ContributionDay `json:"contribution_days"`
}

// ContributionDay represents a single day's contributions
type ContributionDay struct {
	Color             string `json:"color"`
	ContributionCount int    `json:"contribution_count"`
	Date              string `json:"date"`
}

// Repository represents a GitHub repository
type Repository struct {
	Name            string   `json:"name"`
	Description     string   `json:"description"`
	StargazerCount  int      `json:"stargazer_count"`
	ForkCount       int      `json:"fork_count"`
	PrimaryLanguage Language `json:"primary_language"`
	URL             string   `json:"url"`
}

// Language represents a programming language
type Language struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

// UserProfileQuery is the GraphQL query structure for GitHub API
type UserProfileQuery struct {
	User struct {
		Login                   githubv4.String
		Name                    githubv4.String
		AvatarURL               githubv4.String `graphql:"avatarUrl"`
		Bio                     githubv4.String
		ContributionsCollection struct {
			TotalCommitContributions      githubv4.Int
			TotalPullRequestContributions githubv4.Int
			TotalIssueContributions       githubv4.Int
			ContributionCalendar          struct {
				TotalContributions githubv4.Int
				Weeks              []struct {
					ContributionDays []struct {
						Color             githubv4.String
						ContributionCount githubv4.Int
						Date              githubv4.String
					}
				}
			}
		} `graphql:"contributionsCollection(from: $from, to: $to)"`
		PinnedItems struct {
			Nodes []struct {
				Repository struct {
					Name            githubv4.String
					Description     githubv4.String
					StargazerCount  githubv4.Int `graphql:"stargazerCount"`
					ForkCount       githubv4.Int
					PrimaryLanguage struct {
						Name  githubv4.String
						Color githubv4.String
					}
					URL githubv4.String `graphql:"url"`
				} `graphql:"... on Repository"`
			}
		} `graphql:"pinnedItems(first: 6, types: REPOSITORY)"`
		Repositories struct {
			TotalCount githubv4.Int
			Nodes      []struct {
				StargazerCount githubv4.Int `graphql:"stargazerCount"`
			}
		} `graphql:"repositories(first: 100, orderBy: {field: STARGAZERS, direction: DESC}, ownerAffiliations: OWNER, privacy: PUBLIC)"`
	} `graphql:"user(login: $username)"`
}

// FetchUserProfile fetches a GitHub user's profile and contribution statistics
// using the GitHub GraphQL API v4
func FetchUserProfile(username, token string) (*UserProfileStats, error) {
	ctx := context.Background()

	// Create OAuth2 token source
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	httpClient := oauth2.NewClient(ctx, src)

	// Create GitHub GraphQL client
	client := githubv4.NewClient(httpClient)

	// Calculate date range (start of current year to now)
	now := time.Now()
	startOfYear := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, time.UTC)

	// Define query variables
	variables := map[string]interface{}{
		"username": githubv4.String(username),
		"from":     githubv4.DateTime{Time: startOfYear},
		"to":       githubv4.DateTime{Time: now},
	}

	// Execute query
	var query UserProfileQuery
	err := client.Query(ctx, &query, variables)
	if err != nil {
		return nil, err
	}

	// Calculate total stars earned
	totalStarsEarned := 0
	for _, repo := range query.User.Repositories.Nodes {
		totalStarsEarned += int(repo.StargazerCount)
	}

	// Build contribution calendar
	contributionCalendar := ContributionCalendar{
		TotalContributions: int(query.User.ContributionsCollection.ContributionCalendar.TotalContributions),
		Weeks:              make([]ContributionWeek, 0),
	}

	for _, week := range query.User.ContributionsCollection.ContributionCalendar.Weeks {
		contributionWeek := ContributionWeek{
			ContributionDays: make([]ContributionDay, 0),
		}
		for _, day := range week.ContributionDays {
			contributionWeek.ContributionDays = append(contributionWeek.ContributionDays, ContributionDay{
				Color:             string(day.Color),
				ContributionCount: int(day.ContributionCount),
				Date:              string(day.Date),
			})
		}
		contributionCalendar.Weeks = append(contributionCalendar.Weeks, contributionWeek)
	}

	// Build pinned repositories
	pinnedRepos := make([]Repository, 0)
	for _, item := range query.User.PinnedItems.Nodes {
		repo := Repository{
			Name:           string(item.Repository.Name),
			Description:    string(item.Repository.Description),
			StargazerCount: int(item.Repository.StargazerCount),
			ForkCount:      int(item.Repository.ForkCount),
			URL:            string(item.Repository.URL),
			PrimaryLanguage: Language{
				Name:  string(item.Repository.PrimaryLanguage.Name),
				Color: string(item.Repository.PrimaryLanguage.Color),
			},
		}
		pinnedRepos = append(pinnedRepos, repo)
	}

	// Build final stats
	stats := &UserProfileStats{
		Login:                   string(query.User.Login),
		Name:                    string(query.User.Name),
		AvatarURL:               string(query.User.AvatarURL),
		Bio:                     string(query.User.Bio),
		TotalCommits:            int(query.User.ContributionsCollection.TotalCommitContributions),
		TotalPullRequests:       int(query.User.ContributionsCollection.TotalPullRequestContributions),
		TotalIssues:             int(query.User.ContributionsCollection.TotalIssueContributions),
		TotalStarsEarned:        totalStarsEarned,
		ContributionCalendar:    contributionCalendar,
		PinnedRepositories:      pinnedRepos,
		TotalPublicRepositories: int(query.User.Repositories.TotalCount),
	}

	return stats, nil
}

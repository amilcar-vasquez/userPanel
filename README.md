# GitHub Profile Dashboard

An enhanced GitHub profile viewer that displays developer statistics, contribution data, and a calculated developer rank. Built with Go backend and SvelteKit + Material Design 3 frontend.

## Overview

This application fetches and visualizes GitHub profile data through the GitHub GraphQL API, calculating a comprehensive developer rank based on contributions, pull requests, code reviews, stars earned, and followers. Users authenticate with their local account and configure GitHub credentials to view their enhanced profile.

## Features

### GitHub Integration
-  GitHub GraphQL API v4 integration
-  Comprehensive contribution stats (commits, PRs, issues, reviews)
-  Repository stars and follower tracking
-  Contribution calendar visualization
-  Pinned repositories display with language detection

### Developer Rank System
-  7-tier ranking (S+, S, A+, A, B+, B, C)
-  Weighted scoring algorithm (Commits ×2, PRs ×3, Issues ×1, Reviews ×2, Stars ×4, Followers ×1)
-  Progress tracking to next rank
-  Visual rank badges with gradient themes

### Authentication & Profile
-  User registration and JWT authentication
-  Secure GitHub token storage
-  Local profile management (name editing)
-  Account deletion (soft delete)

### UI/UX
-  Material Design 3 with dynamic theming
-  Responsive layout for mobile and desktop
-  Real-time data loading with progress indicators
-  Settings dialog for GitHub credentials

## 🏗️ Architecture

### Backend (Go)
Built with Go 1.24+, using Chi router, GORM ORM, and PostgreSQL. Key components:

- **GitHub Client** (`internal/github/client.go`): GraphQL API integration with oauth2 authentication
- **Rank Calculator** (`internal/github/rank.go`): Weighted scoring algorithm with 7 tier thresholds
- **Auth System**: JWT-based authentication with bcrypt password hashing
- **API Handlers**: RESTful endpoints for user management and GitHub data
- **Middleware**: JWT authentication, CORS, request logging

**GitHub API Integration:**
- Uses `githubv4` library for GraphQL queries
- Fetches user profile, contributions, repositories, followers
- OAuth2 token authentication for secure API access

### Frontend (SvelteKit)
Built with SvelteKit, TypeScript, and Material Web Components 3. Key features:

- **GitHub Profile Page** (`/profile/github`): Main dashboard with stats, rank card, contribution calendar, pinned repos
- **Account Info Page** (`/profile`): Local account management and name editing
- **Auth Pages**: Login and registration with Material 3 form components
- **State Management**: Svelte stores for auth state and user data
- **API Client** (`lib/api.ts`, `lib/github.ts`): Type-safe API communication with error handling

## 📋 Prerequisites

- Docker & Docker Compose
- Go 1.24+ (for local development)
- PostgreSQL 16+ (if running without Docker)

## Start Services

```bash
# Start with Docker Compose
docker compose up --build

# Or in detached mode
docker compose up --build -d

# View logs
docker compose logs -f api
```

The API will be available at `http://localhost:8080`

## API Endpoints

### Health Check

#### `GET /api/health`
Check service status and uptime.

**Response:**
```json
{
  "success": true,
  "data": {
    "status": "ok",
    "uptime": "5m 23s"
  }
}
```

---

### Authentication

#### `POST /api/register`
Register a new user account.

**Request Body:**
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "password123"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "name": "John Doe",
      "email": "john@example.com",
      "created_at": "2025-11-01T17:51:14Z",
      "updated_at": "2025-11-01T17:51:14Z"
    }
  }
}
```
---

```
### Database Schema

The `User` model is auto-migrated on startup:

```go
type User struct {
    ID             uint           // Primary key
    Name           string         // User's full name
    Email          string         // Unique email address
    PasswordHash   string         // bcrypt hashed password
    Avatar         string         // Profile avatar URL (optional, not used)
    GithubUsername string         // GitHub username
    GithubToken    string         // GitHub personal access token
    CreatedAt      time.Time      // Account creation timestamp
    UpdatedAt      time.Time      // Last update timestamp
    DeletedAt      *gorm.DeletedAt // Soft delete timestamp
}
```

### GitHub API Requirements

To use the GitHub integration, users need:
1. A GitHub account
2. A Personal Access Token with permissions:
   - `read:user` - Read user profile data
   - `repo` - Access repository information

Generate a token at: https://github.com/settings/tokens/new

## 🐳 Docker Commands

```bash
# Build and start
docker compose up --build

# Stop services
docker compose down

# View logs
docker compose logs -f api

# Restart API only
docker compose restart api

# Remove volumes (caution: deletes database data)
docker compose down -v
```

## User Flow

1. **Registration/Login** → User creates account with email and password
2. **GitHub Setup** → User clicks "Settings" to configure GitHub username and personal access token
3. **Profile Dashboard** → Application fetches GitHub data and displays:
   - Developer rank with progress bar
   - Contribution statistics (6 stat cards)
   - Yearly contribution calendar
   - Pinned repositories
4. **Account Info** → Users can view local account details and edit their name
5. **Logout/Delete** → Standard account management options

### Port Already in Use
```bash
# Check what's using port 8080
sudo lsof -i :8080

# Or change the port in .env
PORT=8081
```
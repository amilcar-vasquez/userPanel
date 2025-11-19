# GitHub Profile Dashboard

An enhanced GitHub profile viewer that displays developer statistics, contribution data, and a calculated developer rank. Built with Go backend and SvelteKit + Material Design 3 frontend.

## 📊 Overview

This application fetches and visualizes GitHub profile data through the GitHub GraphQL API, calculating a comprehensive developer rank based on contributions, pull requests, code reviews, stars earned, and followers. Users authenticate with their local account and configure GitHub credentials to view their enhanced profile.

## 🚀 Features

### GitHub Integration
- ✅ GitHub GraphQL API v4 integration
- ✅ Comprehensive contribution stats (commits, PRs, issues, reviews)
- ✅ Repository stars and follower tracking
- ✅ Contribution calendar visualization
- ✅ Pinned repositories display with language detection

### Developer Rank System
- ✅ 7-tier ranking (S+, S, A+, A, B+, B, C)
- ✅ Weighted scoring algorithm (Commits ×2, PRs ×3, Issues ×1, Reviews ×2, Stars ×4, Followers ×1)
- ✅ Progress tracking to next rank
- ✅ Visual rank badges with gradient themes

### Authentication & Profile
- ✅ User registration and JWT authentication
- ✅ Secure GitHub token storage
- ✅ Local profile management (name editing)
- ✅ Account deletion (soft delete)

### UI/UX
- ✅ Material Design 3 with dynamic theming
- ✅ Responsive layout for mobile and desktop
- ✅ Real-time data loading with progress indicators
- ✅ Settings dialog for GitHub credentials

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

**Rank Visualization:**
- Gradient rank badges (S+ gold, S pink, A+ cyan, A green, B+ yellow, B silver, C bronze)
- Progress bar showing advancement to next tier
- Dynamic color theming based on rank level

## 🎯 Developer Rank Calculation

**Scoring Formula:**
```
Score = (Commits × 2) + (PRs × 3) + (Issues × 1) + (Reviews × 2) + (Stars × 4) + (Followers × 1)
```

**Rank Tiers:**
- **S+**: 2000+ points (Gold gradient)
- **S**: 1000-1999 points (Pink gradient)
- **A+**: 500-999 points (Cyan gradient)
- **A**: 200-499 points (Green gradient)
- **B+**: 100-199 points (Yellow gradient)
- **B**: 50-99 points (Silver gradient)
- **C**: 0-49 points (Bronze gradient)

## 📋 Prerequisites

- Docker & Docker Compose
- Go 1.24+ (for local development)
- PostgreSQL 16+ (if running without Docker)

## 🛠️ Setup

### 1. Clone and Configure

```bash
cd /path/to/userPanel
```

### 2. Environment Variables

Create or update `backend/.env`:

```env
PORT=8080
DATABASE_URL=postgres://postgres:password@db:5432/authdb?sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-in-production
CORS_ORIGIN=http://localhost:5173
```

### 3. Start Services

```bash
# Start with Docker Compose
docker compose up --build

# Or in detached mode
docker compose up --build -d

# View logs
docker compose logs -f api
```

The API will be available at `http://localhost:8080`

## 📡 API Endpoints

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

**Validation:**
- Email is required and must be unique
- Name is required
- Password must be at least 6 characters

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

**Error Response (409):**
```json
{
  "success": false,
  "message": "Email already registered"
}
```

---

#### `POST /api/login`
Authenticate and receive a JWT token.

**Request Body:**
```json
{
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

**Error Response (401):**
```json
{
  "success": false,
  "message": "Invalid email or password"
}
```

---

### User Profile (Protected Routes)

All profile endpoints require authentication via JWT token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

#### `GET /api/profile`
Get authenticated user's profile.

**Headers:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Success Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "John Doe",
    "email": "john@example.com",
    "avatar": "https://example.com/avatar.jpg",
    "created_at": "2025-11-01T17:51:14Z",
    "updated_at": "2025-11-01T17:51:14Z"
  }
}
```

**Error Response (401):**
```json
{
  "success": false,
  "message": "Invalid or expired token"
}
```

---

#### `PUT /api/profile`
Update authenticated user's profile.

**Headers:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Jane Doe",
  "avatar": "https://example.com/new-avatar.jpg"
}
```

**Success Response (200):**
```json
{
  "success": true,
  "data": {
    "id": 1,
    "name": "Jane Doe",
    "email": "john@example.com",
    "avatar": "https://example.com/new-avatar.jpg",
    "created_at": "2025-11-01T17:51:14Z",
    "updated_at": "2025-11-01T17:52:05Z"
  }
}
```

---

#### `DELETE /api/profile`
Delete authenticated user's account (soft delete).

**Headers:**
```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**Success Response (200):**
```json
{
  "success": true,
  "message": "Account deleted successfully"
}
```

---

## 🧪 Testing with cURL

```bash
# Health check
curl http://localhost:8080/api/health

# Register
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","password":"password123"}'

# Login
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"email":"john@example.com","password":"password123"}'

# Get Profile (replace TOKEN with your JWT)
curl -X GET http://localhost:8080/api/profile \
  -H "Authorization: Bearer <TOKEN>"

# Update Profile
curl -X PUT http://localhost:8080/api/profile \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"name":"Jane Doe","avatar":"https://example.com/avatar.jpg"}'

# Delete Account
curl -X DELETE http://localhost:8080/api/profile \
  -H "Authorization: Bearer <TOKEN>"
```

## 🏗️ Project Structure

```
userPanel/
├── backend/
│   ├── cmd/auth-service/
│   │   └── main.go                 # Application entry point
│   ├── config/
│   │   └── config.go               # Configuration management
│   ├── internal/
│   │   ├── github/
│   │   │   ├── client.go          # GitHub GraphQL API client
│   │   │   └── rank.go            # Developer rank calculator
│   │   ├── handlers/
│   │   │   ├── auth.go            # Auth endpoints (register, login)
│   │   │   ├── user.go            # User profile endpoints
│   │   │   ├── github.go          # GitHub profile endpoints
│   │   │   └── health.go          # Health check endpoint
│   │   ├── middleware/
│   │   │   ├── auth.go            # JWT authentication middleware
│   │   │   └── logger.go          # Request logging middleware
│   │   ├── models/
│   │   │   └── user.go            # User model (with GitHub credentials)
│   │   └── utils/
│   │       ├── jwt.go             # JWT utilities
│   │       ├── password.go        # Password hashing
│   │       └── response.go        # JSON response helpers
│   ├── routes/
│   │   └── routes.go              # Route definitions
│   ├── Dockerfile
│   └── .env
├── frontend/
│   ├── src/
│   │   ├── lib/
│   │   │   ├── api.ts             # API client
│   │   │   ├── github.ts          # GitHub API types & functions
│   │   │   └── stores/
│   │   │       ├── auth.ts        # Auth state management
│   │   │       └── toast.ts       # Toast notifications
│   │   ├── routes/
│   │   │   ├── login/             # Login page
│   │   │   ├── register/          # Registration page
│   │   │   ├── profile/
│   │   │   │   ├── +page.svelte   # Local account info
│   │   │   │   ├── edit/          # Edit name page
│   │   │   │   └── github/        # GitHub profile dashboard
│   │   │   └── +layout.svelte     # Root layout with toast
│   │   └── app.css                # Global styles & Material Web fixes
│   ├── package.json
│   └── svelte.config.js
├── docker-compose.yml             # Multi-service orchestration
└── README.md
```

## 🔐 Security Features

- **Password Hashing**: bcrypt with default cost (10)
- **JWT Tokens**: HS256 algorithm, 24-hour expiry
- **GitHub Token Storage**: Encrypted storage of GitHub personal access tokens
- **CORS**: Configured for frontend origin (default: http://localhost:5173)
- **Soft Deletes**: User accounts are soft-deleted (recoverable)
- **Email Normalization**: Emails are lowercased and trimmed
- **No Sensitive Data Logging**: Passwords and tokens never logged in plaintext

## 🔧 Development

### Local Development (without Docker)

```bash
# Install dependencies
go mod download

# Run database migrations (ensure PostgreSQL is running)
# Update DATABASE_URL in .env to point to your local DB

# Run the server
go run ./backend/cmd/auth-service/main.go
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

## 🌐 User Flow

1. **Registration/Login** → User creates account with email and password
2. **GitHub Setup** → User clicks "Settings" to configure GitHub username and personal access token
3. **Profile Dashboard** → Application fetches GitHub data and displays:
   - Developer rank with progress bar
   - Contribution statistics (6 stat cards)
   - Yearly contribution calendar
   - Pinned repositories
4. **Account Info** → Users can view local account details and edit their name
5. **Logout/Delete** → Standard account management options

## 🎨 Frontend Stack

- **SvelteKit** - Full-stack framework with SSR support
- **TypeScript** - Type-safe development
- **Material Web 3** - Google's Material Design components
- **Vite** - Fast build tool and dev server
- **CSS Custom Properties** - Dynamic theming with Material Design tokens

## 📝 Environment Variables Reference

### Backend (`backend/.env`)

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `PORT` | Server port | `8080` | No |
| `DATABASE_URL` | PostgreSQL connection string | - | Yes |
| `JWT_SECRET` | Secret key for JWT signing | - | Yes |
| `CORS_ORIGIN` | Allowed CORS origin | `http://localhost:5173` | No |

### Services in Docker Compose

- **db**: PostgreSQL 16 database (port 5432)
- **api**: Go backend service (port 8080)
- **frontend**: SvelteKit app with Node.js (port 5173)

## 🤝 Contributing

1. Follow Go best practices and idiomatic code
2. Use descriptive variable names
3. Add comments for complex logic
4. Test endpoints before committing
5. Update this README for new features

## 📄 License

This project is part of the userPanel application.

## 🆘 Troubleshooting

### Port Already in Use
```bash
# Check what's using port 8080
sudo lsof -i :8080

# Or change the port in .env
PORT=8081
```

### Database Connection Issues
- Ensure PostgreSQL container is running: `docker compose ps`
- Check DATABASE_URL in `.env` matches `docker-compose.yml` settings
- Verify network connectivity: `docker network ls`

### JWT Token Invalid
- Ensure JWT_SECRET matches between registration and login
- Check token hasn't expired (24-hour lifetime)
- Verify Authorization header format: `Bearer <token>`

## 🎯 Key Endpoints

### GitHub Integration
- `GET /api/github/profile` - Fetch GitHub profile with calculated rank
- `PUT /api/github/credentials` - Save GitHub username and token

### Authentication
- `POST /api/register` - Create new account
- `POST /api/login` - Authenticate and get JWT token

### User Profile
- `GET /api/profile` - Get user account info
- `PUT /api/profile` - Update user name
- `DELETE /api/profile` - Delete account (soft delete)

### Health
- `GET /api/health` - Service health check

---

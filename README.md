# Auth Service API

A RESTful authentication service built with Go, PostgreSQL, JWT, and Docker. Designed for integration with a SvelteKit + Material 3 frontend.

## 🚀 Features

- ✅ User registration and authentication
- ✅ JWT-based token authentication (24-hour expiry)
- ✅ Secure password hashing with bcrypt
- ✅ CORS enabled for frontend integration
- ✅ RESTful API with consistent JSON responses
- ✅ Request logging middleware
- ✅ PostgreSQL with GORM ORM
- ✅ Docker & Docker Compose ready
- ✅ Auto database migrations

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
backend/
├── cmd/
│   └── auth-service/
│       └── main.go              # Application entry point
├── config/
│   └── config.go                # Configuration management
├── internal/
│   ├── handlers/
│   │   ├── auth.go             # Auth endpoints (register, login)
│   │   ├── user.go             # User profile endpoints
│   │   └── health.go           # Health check endpoint
│   ├── middleware/
│   │   ├── auth.go             # JWT authentication middleware
│   │   └── logger.go           # Request logging middleware
│   ├── models/
│   │   └── user.go             # User database model
│   └── utils/
│       ├── jwt.go              # JWT utilities
│       ├── password.go         # Password hashing
│       └── response.go         # JSON response helpers
├── routes/
│   └── routes.go               # Route definitions
├── db/
│   ├── migrations/             # SQL migrations (optional)
│   └── queries/                # SQL queries (if using sqlc)
├── Dockerfile
├── .env                        # Environment variables
└── .env.example                # Example environment file
```

## 🔐 Security Features

- **Password Hashing**: bcrypt with default cost (10)
- **JWT Tokens**: HS256 algorithm, 24-hour expiry
- **CORS**: Configured for frontend origin (default: http://localhost:5173)
- **Soft Deletes**: User accounts are soft-deleted (recoverable)
- **Email Normalization**: Emails are lowercased and trimmed
- **No Password Logging**: Passwords are never logged in plaintext

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
    ID           uint           // Primary key
    Name         string         // User's full name
    Email        string         // Unique email address
    PasswordHash string         // bcrypt hashed password
    Avatar       string         // Profile avatar URL (optional)
    CreatedAt    time.Time      // Account creation timestamp
    UpdatedAt    time.Time      // Last update timestamp
    DeletedAt    *time.Time     // Soft delete timestamp
}
```

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

## 🌐 Frontend Integration

### SvelteKit Example

```javascript
// Login function
async function login(email, password) {
  const response = await fetch('http://localhost:8080/api/login', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ email, password })
  });
  
  const data = await response.json();
  if (data.success) {
    localStorage.setItem('token', data.data.token);
    return data.data.user;
  }
  throw new Error(data.message);
}

// Authenticated request
async function getProfile() {
  const token = localStorage.getItem('token');
  const response = await fetch('http://localhost:8080/api/profile', {
    headers: { 'Authorization': `Bearer ${token}` }
  });
  
  const data = await response.json();
  if (data.success) return data.data;
  throw new Error(data.message);
}
```

## 📝 Environment Variables Reference

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `PORT` | Server port | `8080` | No |
| `DATABASE_URL` | PostgreSQL connection string | - | Yes |
| `JWT_SECRET` | Secret key for JWT signing | - | Yes |
| `CORS_ORIGIN` | Allowed CORS origin | `http://localhost:5173` | No |

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

---

**Built with** ❤️ **using Go, PostgreSQL, GORM, Chi, and Docker**

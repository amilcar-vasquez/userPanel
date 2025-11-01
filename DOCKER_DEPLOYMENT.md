# Docker Deployment Guide

## Overview

The Auth Service is fully containerized and can be deployed using Docker Compose. The setup includes:
- **PostgreSQL Database** (port 5432 - internal only)
- **Go Backend API** (port 8080)
- **SvelteKit Frontend** (port 3000)

## Prerequisites

- Docker Engine 20.10+
- Docker Compose 2.0+

## Quick Start

### 1. Build and Start All Services

```bash
docker compose up --build -d
```

This command will:
- Build the backend API container
- Build the frontend container
- Start the PostgreSQL database
- Start all services in detached mode

### 2. Verify Services

Check that all services are running:
```bash
docker compose ps
```

Expected output:
```
NAME                   IMAGE                COMMAND                  SERVICE    CREATED          STATUS          PORTS
userpanel-api-1        userpanel-api        "./main"                 api        10 seconds ago   Up 8 seconds    0.0.0.0:8080->8080/tcp
userpanel-db-1         postgres:16          "docker-entrypoint.s…"   db         10 seconds ago   Up 9 seconds    5432/tcp
userpanel-frontend-1   userpanel-frontend   "docker-entrypoint.s…"   frontend   10 seconds ago   Up 8 seconds    0.0.0.0:3000->3000/tcp
```

### 3. Access the Application

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080/api
- **Health Check**: http://localhost:8080/api/health

## Configuration

### Backend Environment Variables

Located in `backend/.env`:
```env
PORT=8080
DATABASE_URL=postgres://postgres:password@db:5432/authdb?sslmode=disable
JWT_SECRET=your-super-secret-jwt-key-change-this-in-production
CORS_ORIGIN=http://localhost:3000
```

**Note**: In Docker, the database host is `db` (the service name), not `localhost`.

### Frontend Environment Variables

Located in `frontend/.env`:
```env
PUBLIC_API_URL=http://localhost:8080/api
```

For production, update this to your actual API URL.

## Docker Commands

### View Logs

View logs from all services:
```bash
docker compose logs -f
```

View logs from a specific service:
```bash
docker compose logs -f api
docker compose logs -f frontend
docker compose logs -f db
```

### Stop Services

```bash
docker compose down
```

### Stop and Remove Volumes (⚠️ deletes database data)

```bash
docker compose down -v
```

### Rebuild a Specific Service

```bash
docker compose up --build api
docker compose up --build frontend
```

### Execute Commands in Containers

Access the backend container:
```bash
docker compose exec api sh
```

Access the frontend container:
```bash
docker compose exec frontend sh
```

Access the database:
```bash
docker compose exec db psql -U postgres -d authdb
```

## Development vs Production

### Development Mode

For development, you can still run services locally:

1. Start only the database:
   ```bash
   docker compose up db -d
   ```

2. Run backend locally:
   ```bash
   cd backend
   go run cmd/auth-service/main.go
   ```

3. Run frontend locally:
   ```bash
   cd frontend
   npm run dev
   ```

### Production Deployment

For production:

1. Update environment variables in `backend/.env`:
   - Use a strong JWT secret
   - Update database credentials
   - Set proper CORS origin

2. Update `frontend/.env`:
   - Set `PUBLIC_API_URL` to your production API URL

3. Consider using a reverse proxy (nginx) in front of both services

4. Use Docker secrets or environment variables for sensitive data

## Troubleshooting

### Port Already in Use

If you get a port conflict error:

1. Check what's using the port:
   ```bash
   sudo lsof -i :8080
   sudo lsof -i :3000
   sudo lsof -i :5432
   ```

2. Either stop the conflicting service or change the port in `docker-compose.yml`

### Database Connection Issues

If the backend can't connect to the database:

1. Make sure the database host is set to `db` (not `localhost`) in `backend/.env`
2. Check database logs: `docker compose logs db`
3. Restart the API service: `docker compose restart api`

### Frontend Can't Connect to Backend

1. Check that the API is running: `curl http://localhost:8080/api/health`
2. Verify `PUBLIC_API_URL` in `frontend/.env`
3. Check browser console for CORS errors
4. Ensure `CORS_ORIGIN` in backend matches your frontend URL

### Rebuild from Scratch

If you encounter persistent issues:

```bash
# Stop and remove everything
docker compose down -v

# Remove images
docker compose rm -f

# Rebuild
docker compose up --build
```

## Architecture

```
┌─────────────────┐
│   Browser       │
└────────┬────────┘
         │
         │ HTTP (port 3000)
         │
         ▼
┌─────────────────┐      HTTP (port 8080)      ┌─────────────────┐
│   Frontend      │─────────────────────────────▶│   Backend API   │
│   (SvelteKit)   │                              │   (Go)          │
└─────────────────┘                              └────────┬────────┘
                                                          │
                                                          │ PostgreSQL
                                                          │
                                                          ▼
                                                 ┌─────────────────┐
                                                 │   Database      │
                                                 │   (PostgreSQL)  │
                                                 └─────────────────┘
```

## Security Notes

⚠️ **Important**: The default configuration is for development only.

For production:
- Change all default passwords
- Use a strong, random JWT secret (32+ characters)
- Enable HTTPS/TLS
- Use environment-specific secrets management
- Review and harden CORS settings
- Enable database SSL connections
- Regular security updates for all dependencies

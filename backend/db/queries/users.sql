-- Queries for user management with GitHub integration support

-- 1. Create a new user
INSERT INTO users (name, email, password_hash, created_at, updated_at)
VALUES ($1, $2, $3, NOW(), NOW())
RETURNING id, name, email, created_at, updated_at;

-- 2. Get user by email (for login)
SELECT id, name, email, password_hash, avatar, github_username, github_token, created_at, updated_at
FROM users
WHERE email = $1 AND deleted_at IS NULL;

-- 3. Get user by id (for profile)
SELECT id, name, email, avatar, github_username, created_at, updated_at
FROM users
WHERE id = $1 AND deleted_at IS NULL;

-- 4. Update user profile (name and avatar)
UPDATE users
SET name = $1, avatar = $2, updated_at = NOW()
WHERE id = $3 AND deleted_at IS NULL
RETURNING id, name, email, avatar, github_username, created_at, updated_at;

-- 5. Update GitHub credentials
UPDATE users
SET github_username = $1, github_token = $2, updated_at = NOW()
WHERE id = $3 AND deleted_at IS NULL
RETURNING id, name, email, github_username, created_at, updated_at;

-- 6. Soft delete user
UPDATE users
SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL
RETURNING id;

-- 7. List all active users (admin query)
SELECT id, name, email, github_username, created_at, updated_at
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC;
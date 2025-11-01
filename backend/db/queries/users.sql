-- Queries for user management
-- 1. Create a new user
INSERT INTO users (email, password_hash, created_at)
VALUES ($1, $2, NOW())
RETURNING id, email, created_at;

-- 2. Get user by email
SELECT id, email, password_hash, created_at, updated_at
FROM users
WHERE email = $1;

-- 3. Update user password
UPDATE users
SET password_hash = $1, updated_at = NOW()
WHERE id = $2
RETURNING id, email, updated_at;

-- 4. Delete user by id
DELETE FROM users
WHERE id = $1
RETURNING id;
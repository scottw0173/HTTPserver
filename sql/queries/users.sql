-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: DeleteAllUsers :exec
DELETE FROM users;

-- name: CreateChirp :one
INSERT INTO chirps (id, created_at, updated_at, body, user_id)
VALUEs (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: ReturnChirps :many
SELECT * 
FROM chirps
ORDER BY created_at ASC;

-- name: ReturnSingleChirp :one
SELECT *
FROM chirps
WHERE id = $1;

-- name: ReturnUser :one
SELECT *
FROM users
WHERE email = $1;
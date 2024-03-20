-- name: getAllUsers :many
SELECT * FROM users;

-- name: getUserByUsername :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE username = $1;

-- name: getUserByID :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE user_id = $1;

-- name: getUserByEmail :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE user_email = $1;

-- name: createUser :one
INSERT INTO users (username, user_email, user_password, first_name, last_name)
VALUES ($1, $2, $3, $4, $5)
RETURNING user_id;
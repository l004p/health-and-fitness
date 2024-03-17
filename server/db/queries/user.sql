-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByUsername :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE username = $1;

-- name: GetUserByID :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE user_id = $1;

-- name: GetUserByEmail :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE user_email = $1;
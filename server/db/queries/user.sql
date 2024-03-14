-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByUsername :one
SELECT user_id, first_name, last_name
FROM users
WHERE username = $1;
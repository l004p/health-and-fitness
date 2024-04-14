-- name: getAllUsers :many
SELECT * FROM users;

-- name: getUserByUsername :one
SELECT *
FROM users
WHERE username = $1;

-- name: getUserByID :one
SELECT user_id, user_email, username, first_name, last_name
FROM users
WHERE user_id = $1;

-- name: getUserByEmail :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE user_email = $1;

-- name: createUser :one
INSERT INTO users (username, user_email, user_password, first_name, last_name)
VALUES ($1, $2, $3, $4, $5)
RETURNING user_id, username, first_name, last_name, user_email;

-- name: getUserRoles :many
SELECT r.role_name
FROM user_roles ur
INNER JOIN users u 
ON u.user_id=ur.user_id
INNER JOIN roles r 
ON r.role_id=ur.role_id
WHERE u.user_id=$1;

-- name: verifyUserRole :one
SELECT u.user_id
FROM user_roles ur
INNER JOIN users u 
ON u.user_id=ur.user_id
INNER JOIN roles r 
ON r.role_id=ur.role_id
WHERE u.user_id=$1 AND r.role_name=$2;


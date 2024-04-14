// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package pg

import (
	"context"
)

const createUser = `-- name: createUser :one
INSERT INTO users (username, user_email, user_password, first_name, last_name)
VALUES ($1, $2, $3, $4, $5)
RETURNING user_id, username, first_name, last_name, user_email
`

type createUserParams struct {
	Username     string
	UserEmail    string
	UserPassword string
	FirstName    string
	LastName     string
}

type createUserRow struct {
	UserID    int32
	Username  string
	FirstName string
	LastName  string
	UserEmail string
}

func (q *Queries) createUser(ctx context.Context, arg createUserParams) (createUserRow, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.UserEmail,
		arg.UserPassword,
		arg.FirstName,
		arg.LastName,
	)
	var i createUserRow
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.UserEmail,
	)
	return i, err
}

const getAllUsers = `-- name: getAllUsers :many
SELECT user_id, username, user_password, first_name, last_name, user_email FROM users
`

func (q *Queries) getAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.Username,
			&i.UserPassword,
			&i.FirstName,
			&i.LastName,
			&i.UserEmail,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: getUserByEmail :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE user_email = $1
`

type getUserByEmailRow struct {
	UserID    int32
	UserEmail string
	FirstName string
	LastName  string
}

func (q *Queries) getUserByEmail(ctx context.Context, userEmail string) (getUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, userEmail)
	var i getUserByEmailRow
	err := row.Scan(
		&i.UserID,
		&i.UserEmail,
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}

const getUserByID = `-- name: getUserByID :one
SELECT user_id, user_email, username, first_name, last_name
FROM users
WHERE user_id = $1
`

type getUserByIDRow struct {
	UserID    int32
	UserEmail string
	Username  string
	FirstName string
	LastName  string
}

func (q *Queries) getUserByID(ctx context.Context, userID int32) (getUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, userID)
	var i getUserByIDRow
	err := row.Scan(
		&i.UserID,
		&i.UserEmail,
		&i.Username,
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}

const getUserByUsername = `-- name: getUserByUsername :one
SELECT user_id, username, user_password, first_name, last_name, user_email
FROM users
WHERE username = $1
`

func (q *Queries) getUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.Username,
		&i.UserPassword,
		&i.FirstName,
		&i.LastName,
		&i.UserEmail,
	)
	return i, err
}

const getUserRoles = `-- name: getUserRoles :many
SELECT r.role_name
FROM user_roles ur
INNER JOIN users u 
ON u.user_id=ur.user_id
INNER JOIN roles r 
ON r.role_id=ur.role_id
WHERE u.user_id=$1
`

func (q *Queries) getUserRoles(ctx context.Context, userID int32) ([]string, error) {
	rows, err := q.db.Query(ctx, getUserRoles, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var role_name string
		if err := rows.Scan(&role_name); err != nil {
			return nil, err
		}
		items = append(items, role_name)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const verifyUserRole = `-- name: verifyUserRole :one
SELECT u.user_id
FROM user_roles ur
INNER JOIN users u 
ON u.user_id=ur.user_id
INNER JOIN roles r 
ON r.role_id=ur.role_id
WHERE u.user_id=$1 AND r.role_name=$2
`

type verifyUserRoleParams struct {
	UserID   int32
	RoleName string
}

func (q *Queries) verifyUserRole(ctx context.Context, arg verifyUserRoleParams) (int32, error) {
	row := q.db.QueryRow(ctx, verifyUserRole, arg.UserID, arg.RoleName)
	var user_id int32
	err := row.Scan(&user_id)
	return user_id, err
}

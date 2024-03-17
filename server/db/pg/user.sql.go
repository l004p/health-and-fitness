// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package pg

import (
	"context"
)

const getAllUsers = `-- name: GetAllUsers :many
SELECT user_id, username, user_password, first_name, last_name, user_email FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
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

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE user_email = $1
`

type GetUserByEmailRow struct {
	UserID    int32
	UserEmail string
	FirstName string
	LastName  string
}

func (q *Queries) GetUserByEmail(ctx context.Context, userEmail string) (GetUserByEmailRow, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, userEmail)
	var i GetUserByEmailRow
	err := row.Scan(
		&i.UserID,
		&i.UserEmail,
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE user_id = $1
`

type GetUserByIDRow struct {
	UserID    int32
	UserEmail string
	FirstName string
	LastName  string
}

func (q *Queries) GetUserByID(ctx context.Context, userID int32) (GetUserByIDRow, error) {
	row := q.db.QueryRow(ctx, getUserByID, userID)
	var i GetUserByIDRow
	err := row.Scan(
		&i.UserID,
		&i.UserEmail,
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT user_id, user_email, first_name, last_name
FROM users
WHERE username = $1
`

type GetUserByUsernameRow struct {
	UserID    int32
	UserEmail string
	FirstName string
	LastName  string
}

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (GetUserByUsernameRow, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i GetUserByUsernameRow
	err := row.Scan(
		&i.UserID,
		&i.UserEmail,
		&i.FirstName,
		&i.LastName,
	)
	return i, err
}

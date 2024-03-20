package duser

import "context"

type User struct {
	UserID       int32
	Username     string
	UserPassword string
	FirstName    string
	LastName     string
	UserEmail    string
}

type UserRepository interface {
	TestFunction(ctx context.Context) int
	CreateUser(ctx context.Context, input User) (int32, error)
}
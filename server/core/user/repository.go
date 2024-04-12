package user

import "context"

type User struct {
	UserID       int32
	Username     string
	UserPassword string
	FirstName    string
	LastName     string
	UserEmail    string
}

type RoleName string

const (
	MEMBER RoleName = "MEMBER"
	TRAINER RoleName = "TRAINER"
	ADMIN RoleName = "ADMIN"
)

type UserRepository interface {
	TestFunction(ctx context.Context) int
	CreateUser(ctx context.Context, input User) (User, error)
	UserByEmail(ctx context.Context, email string) (User, error)
	UserByID(ctx context.Context, id string) (User, error)
	UserByUsername(ctx context.Context, username string) (User, error)
	HasRole(ctx context.Context, id string, role string) (bool, error)
	FindOrCreate(ctx context.Context, input User) (error)
}
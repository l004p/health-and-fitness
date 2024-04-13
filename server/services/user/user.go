package userservice

import (
	"context"
	//"fmt"
	//"server/db/pg"
	"server/core"

	"golang.org/x/crypto/bcrypt"
)


func TestFunction (r core.UserRepository, ctx context.Context) int {
	return r.TestFunction(ctx)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func CreateUser(r core.UserRepository, ctx context.Context, input core.User) (core.User, error){
	pw, err := HashPassword(input.UserPassword)
	if err != nil {
		return core.User{}, err
	}
	input.UserPassword = pw
	createdUser, err := r.CreateUser(ctx, input)
	if err != nil {
		return core.User{}, err
	}
	return createdUser, nil
}

func AllUsers(r core.UserRepository, ctx context.Context) ([]core.User, error){
	var users []core.User


	return users, nil
}

func UserByEmail(r core.UserRepository, ctx context.Context, email string) (core.User, error) {
	u, err := r.UserByEmail(ctx, email)
	if err != nil {
		return core.User{}, err
	}
	return u, nil
}

func UserByID(r core.UserRepository, ctx context.Context, userID string) (core.User, error) {
	u, err := r.UserByID(ctx, userID)
	if err != nil {
		return core.User{}, err
	}
	return u, nil
}

func UserByUsername(r core.UserRepository, ctx context.Context, username string) (core.User, error) {
	u, err := r.UserByID(ctx, username)
	if err != nil {
		return core.User{}, err
	}
	return u, nil
}

func RoleByID(ctx context.Context, userID string) (string, error) {
	return userID, nil
}

func HasRole(r core.UserRepository, ctx context.Context, id string, role string) (bool, error) {
	result, err := r.HasRole(ctx, id, role)
	if err != nil {
		return false, err
	}
	return result, nil
}
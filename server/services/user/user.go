package userservice

import (
	"context"
	//"fmt"
	//"server/db/pg"
	"server/core/user"

	"golang.org/x/crypto/bcrypt"
)


func TestFunction (r user.UserRepository, ctx context.Context) int {
	return r.TestFunction(ctx)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func CreateUser(r user.UserRepository, ctx context.Context, input user.User) (user.User, error){
	pw, err := HashPassword(input.UserPassword)
	if err != nil {
		return user.User{}, err
	}
	input.UserPassword = pw
	createdUser, err := r.CreateUser(ctx, input)
	if err != nil {
		return user.User{}, err
	}
	return createdUser, nil
}

func AllUsers(r user.UserRepository, ctx context.Context) ([]user.User, error){
	var users []user.User


	return users, nil
}

func UserByEmail(r user.UserRepository, ctx context.Context, email string) (user.User, error) {
	u, err := r.UserByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func UserByID(r user.UserRepository, ctx context.Context, userID string) (user.User, error) {
	u, err := r.UserByID(ctx, userID)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func UserByUsername(r user.UserRepository, ctx context.Context, username string) (user.User, error) {
	u, err := r.UserByID(ctx, username)
	if err != nil {
		return user.User{}, err
	}
	return u, nil
}

func RoleByID(ctx context.Context, userID string) (string, error) {
	return userID, nil
}

func HasRole(r user.UserRepository, ctx context.Context, id string, role string) (bool, error) {
	result, err := r.HasRole(ctx, id, role)
	if err != nil {
		return false, err
	}
	return result, nil
}
package userservice

import (
	"context"
	//"fmt"
	"server/db/pg"
	"server/core/user"

	"golang.org/x/crypto/bcrypt"
)


func TestFunction (r duser.UserRepository, ctx context.Context) int {
	return r.TestFunction(ctx)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func CreateUser(r duser.UserRepository, ctx context.Context, input duser.User) (int32, error){
	var id int32
	pw, err := HashPassword(input.UserPassword)
	if err != nil {
		return -1, err
	}
	input.UserPassword = pw
	id, err = r.CreateUser(ctx, input)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func AllUsers(ctx context.Context) ([]pg.User, error){
	var users []pg.User

	return users, nil
}

func UserByEmail(ctx context.Context, email string) (pg.User, error) {
	var user pg.User

	return user, nil
}

func UserByID(ctx context.Context, userID string) (pg.User, error) {
	var user pg.User

	return user, nil
}

func UserByUsername(ctx context.Context, username string) (pg.User, error) {
	var user pg.User

	return user, nil
}

func RoleByID(ctx context.Context, userID string) (string, error) {
	return userID, nil
}
package user

import (
	"context"
	"server/db/pg"

	"golang.org/x/crypto/bcrypt"
)


func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
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
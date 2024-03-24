package authentication

import (
	// "os"
	//"fmt"
	"server/core/user"
	"server/services/user"

	//"server/services/user"
	"context"
	"strconv"
	// "github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
)

func Login(r user.UserRepository, ctx context.Context, username string, password string) (string, error) {
	u, err := r.UserByUsername(ctx, username)
	if err != nil {
		return "", err
	}
	err = userservice.CheckPasswordHash(u.UserPassword, password)
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(strconv.FormatInt(int64(u.UserID), 10))
	if err != nil {
		return "", err
	}
	return token, nil
}

func Register(r user.UserRepository, ctx context.Context, input user.User) error {
	err := r.FindOrCreate(ctx, input)
	if err != nil {
		return err
	}
	return nil
}
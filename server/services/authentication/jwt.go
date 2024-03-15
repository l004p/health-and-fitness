package authentication

import (
	"os"
	//"server/services/user"
	"time"
	"fmt"
	//"log"

	"github.com/dgrijalva/jwt-go"
	//"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	UserRole string `json:"user_role"`
	jwt.StandardClaims
}

func GenerateToken(userID string, userRole string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwtCustomClaims{
		UserID: userID,
		UserRole: userRole,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwtCustomClaims{}, func (t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("issue with signing method")
		}
		return jwtSecret, nil
	})
}
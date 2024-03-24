package rest

import (
	//"net/http"
	//"encoding/json"
	//"fmt"
	"server/core/user"
	//"server/services/authentication"
)

type AuthHandler struct {
	ur user.UserRepository
}

func NewAuthHandler(ur user.UserRepository) *AuthHandler {
	return &AuthHandler{ ur: ur }
}
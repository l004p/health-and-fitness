package rest

import (
	//"net/http"
	//"encoding/json"
	//"fmt"
	"server/core"
	//"server/services/authentication"
)

type AuthHandler struct {
	ur core.UserRepository
}

func NewAuthHandler(ur core.UserRepository) *AuthHandler {
	return &AuthHandler{ ur: ur }
}
package middleware

import(
	"server/core/user"
)

type MiddlewareHandler struct {
	ur user.UserRepository
}

func NewMiddlewareHandler(ur user.UserRepository) *MiddlewareHandler {
	return &MiddlewareHandler{ ur: ur }
}
package middleware

import(
	"server/core"
)

type MiddlewareHandler struct {
	ur core.UserRepository
}

func NewMiddlewareHandler(ur core.UserRepository) *MiddlewareHandler {
	return &MiddlewareHandler{ ur: ur }
}
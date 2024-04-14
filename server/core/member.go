package core

import (
	"context"
	//"server/core/trainer"
)

// type Member struct {
// 	user user.User
// }

type MemberRepository interface {
	GetAllTrainers(ctx context.Context, memberID string) ([]User, error)
}
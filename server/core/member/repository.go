package member

import (
	"context"
	//"server/core/trainer"
	"server/core/user"
)

// type Member struct {
// 	user user.User
// }

type MemberRepository interface {
	GetAllTrainers(ctx context.Context, member user.User, trainer user.User) ([]user.User, error)
}
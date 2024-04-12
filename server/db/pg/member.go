package pg

import (
	"context"
	"server/core/user"
	//"server/core/member"
	//"server/core/trainer"
	//"server/core/membership"
)

func (r *repoService) GetAllTrainers(ctx context.Context, member user.User, trainer user.User) ([]user.User, error) {
	panic("not implemented")
}
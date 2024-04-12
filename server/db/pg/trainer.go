package pg

import (
	"context"
	"server/core/user"
	//"server/core/member"
	"server/core/trainer"
	//"server/core/membership"
)

func (r *repoService) AddMemberToRoster(ctx context.Context, trainer user.User, member user.User) (error) {
	panic("not implemented")
}

func (r *repoService) RemoveMemberFromRoster(ctx context.Context, trainer user.User, member user.User) (error) {
	panic("not implemented")
}

func (r *repoService) GetAllMembersTraining(ctx context.Context, trainer user.User, member user.User) ([]user.User, error) {
	panic("not implemented")
}

func (r *repoService) AddInterest(ctx context.Context, trainer user.User, interest trainer.Interest) error {
	panic("not implemented")
}
package pg

import (
	"context"
	"server/core"
	//"server/core/member"
	
	//"server/core/membership"
)

func (r *repoService) AddMemberToRoster(ctx context.Context, trainer core.User, member core.User) (error) {
	panic("not implemented")
}

func (r *repoService) RemoveMemberFromRoster(ctx context.Context, trainer core.User, member core.User) (error) {
	panic("not implemented")
}

func (r *repoService) GetAllMembersTraining(ctx context.Context, trainer core.User, member core.User) ([]core.User, error) {
	panic("not implemented")
}

func (r *repoService) AddInterest(ctx context.Context, trainer core.User, interest core.Interest) error {
	panic("not implemented")
}
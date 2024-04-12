package pg

import (
	"context"
	"server/core/user"
	//"server/core/member"
	//"server/core/trainer"
	"server/core/membership"
)

func (r *repoService) AddMembership(ctx context.Context, user user.User, membership membership.MembershipType) (error) {
	panic("not implemented")
}

func (r *repoService) CancelMembership(ctx context.Context, user user.User, membership membership.MembershipType) (error){
	panic("not implemented")
}

func (r *repoService) ValidMembership(ctx context.Context, user user.User) (bool, error) {
	panic("not implemented")
}
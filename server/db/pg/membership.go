package pg

import (
	"context"
	"server/core"
	//"server/core/member"
	//"server/core/trainer"
)

func (r *repoService) AddMembership(ctx context.Context, user core.User, membership core.MembershipType) (error) {
	panic("not implemented")
}

func (r *repoService) CancelMembership(ctx context.Context, user core.User, membership core.MembershipType) (error){
	panic("not implemented")
}

func (r *repoService) ValidMembership(ctx context.Context, user core.User) (bool, error) {
	panic("not implemented")
}
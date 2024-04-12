package trainer

import (
	"context"
	"server/core/user"
	//"server/core/member"
)

// type Trainer struct {
// 	user user.User
// }

type Interest string

const (
	CARDIO Interest = "CARDIO"
	WEIGHTLOSS Interest = "WEIGHTLOSS"
	STRENGTHTRAINING Interest = "STRENGTHTRAINING"
	ENDURANCE Interest = "ENDURANCE"
	CYCLING Interest = "CYCLING"
	RUNNING Interest = "RUNNING"
	MOBILITY Interest = "MOBILITY"
)

type TrainerRepository interface {
	AddMemberToRoster(ctx context.Context, trainer user.User, member user.User) (error)
	RemoveMemberFromRoster(ctx context.Context, trainer user.User, member user.User) (error)
	GetAllMembersTraining(ctx context.Context, trainer user.User, member user.User) ([]user.User, error)
	AddInterest(ctx context.Context, trainer user.User, interest Interest) (error)
	
}
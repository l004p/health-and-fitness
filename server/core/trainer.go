package core

import (
	"context"
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
	AddMemberToRoster(ctx context.Context, trainer User, member User) (error)
	RemoveMemberFromRoster(ctx context.Context, trainer User, member User) (error)
	GetAllMembersTraining(ctx context.Context, trainer User, member User) ([]User, error)
	AddInterest(ctx context.Context, trainer User, interest Interest) (error)
	SearchForUser(ctx context.Context, firstname string, lastname string) ([]User, error)
}
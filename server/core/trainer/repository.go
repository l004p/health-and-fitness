package trainer

import (
	"context"
	"server/core/user"
	//"server/core/member"
)

// type Trainer struct {
// 	user user.User
// }

type TrainerRepository interface {
	AddMemberToRoster(ctx context.Context, trainer user.User, member user.User) (error);
	RemoveMemberFromRoster(ctx context.Context, trainer user.User, member user.User) (error);
	GetAllMembersTraining(ctx context.Context, trainer user.User, member user.User) ([]user.User, error);
}
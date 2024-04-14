package memberservice 

import(
	"server/core"
	"context"
)

func GetTrainers (r core.MemberRepository, ctx context.Context, memberID string) ([]core.User, error) {
	t, err := r.GetAllTrainers(ctx, memberID)
	if err != nil {
		return nil, err
	}
	return t, nil
}
package pg

import (
	"context"
	"server/core"
	"strconv"
	//"server/core/member"
	//"server/core/trainer"
	//"server/core/membership"
)

func (r *repoService) GetAllTrainers(ctx context.Context, memberID string) ([]core.User, error) {
	//panic("not implemented")
	idInt, err := strconv.ParseInt(memberID, 10, 32)
	if err != nil {
		return nil, err
	}
	t, err := r.getAllTrainers(ctx, int32(idInt))
	if err != nil {
		return nil, err
	}
	var result []core.User
	for _, trainer := range t {
		curr := core.User{
			FirstName: trainer.FirstName,
			LastName: trainer.LastName,
			Username: trainer.Username,
		}
		result = append(result, curr)
	}
	return result, nil
}
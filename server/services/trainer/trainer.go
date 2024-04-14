package trainerservice

import(
	"server/core"
	"context"
) 

func SearchForUser(r core.TrainerRepository, ctx context.Context, firstname string, lastname string) ([]core.User, error) {
	u, err := r.SearchForUser(ctx, firstname, lastname)
	if err != nil {
		return nil, err
	}
	return u, nil
}
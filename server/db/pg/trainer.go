package pg

import (
	"context"
	//"fmt"
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

func (r *repoService) SearchForUser(ctx context.Context, firstname string, lastname string) ([]core.User, error) {
	//panic("not implemented")
	params := getUserByNameParams{
		FirstName: firstname,
		LastName: lastname,
	}
	u, err := r.getUserByName(ctx, params)
	var result []core.User
	if err != nil {
		return nil, err
	}
	//fmt.Println(firstname)
	for _, current := range u {
		//fmt.Println(current.FirstName)
		user := core.User{
			Username: current.Username,
			FirstName: current.FirstName,
			LastName: current.LastName,
		}
		result = append(result, user)
	}
	return result, nil
}
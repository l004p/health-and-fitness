package pg

import(
	"context"
	"server/core"
	//"server/db/connect"
	"github.com/jackc/pgx/v5"
	//"fmt"
	//"log"
	"strconv"
)



func (r *repoService) CreateUser(ctx context.Context, input core.User) (core.User, error) {
	insert := createUserParams{
		Username: input.Username,
		UserEmail: input.UserEmail,
		UserPassword: input.UserPassword,
		FirstName: input.FirstName,
		LastName: input.LastName,
	}
	result, err := r.createUser(ctx, insert)
	if err != nil {
		return core.User{}, err
	}
	createdUser := core.User{
		UserID: result.UserID,
		Username: result.Username,
		UserEmail: result.UserEmail,
		FirstName: result.FirstName,
		LastName: result.LastName,
	}
	return createdUser, nil

}

func (r *repoService) UserByEmail(ctx context.Context, email string) (core.User, error) {
	result, err := r.getUserByEmail(ctx, email)
	if err != nil {
		return core.User{}, err
	}
	user := core.User{
		UserID: result.UserID,
		UserEmail: result.UserEmail,
		FirstName: result.FirstName,
		LastName: result.LastName,
	}
	return user, nil
}

func (r *repoService) UserByID(ctx context.Context, id string) (core.User, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return core.User{}, err
	}
	result, err := r.getUserByID(ctx, int32(idInt))
	if err != nil {
		return core.User{}, err
	}
	user := core.User{
		UserID: result.UserID,
		UserEmail: result.UserEmail,
		Username: result.Username,
		FirstName: result.FirstName,
		LastName: result.LastName,
	}
	return user, nil
}

func (r *repoService) UserByUsername(ctx context.Context, username string) (core.User, error) {
	result, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return core.User{}, err
	}
	user := core.User{
		UserID: result.UserID,
		UserPassword: result.UserPassword,
		UserEmail: result.UserEmail,
		FirstName: result.FirstName,
		LastName: result.LastName,
	}
	return user, nil
}

func (r *repoService) HasRole(ctx context.Context, id string, role string) (bool, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return false, err
	}
	input := verifyUserRoleParams{
		UserID: int32(idInt),
		RoleName: role,
	}
	_, err = r.verifyUserRole(ctx, input)
	if err == pgx.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

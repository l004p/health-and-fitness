package pg

import (
	"context"
	"server/core/user"
	"server/db/connect"
	"github.com/jackc/pgx/v5"
	"fmt"
	"log"
	"strconv"
)

type Repository interface {
	user.UserRepository
}

type repoService struct {
	*Queries 
	db *connect.Postgres
}

func NewRepository(db *connect.Postgres) Repository {
	return &repoService{
		Queries: New(db.Db),
		db: db,
	}
}

func (r *repoService) TestFunction(ctx context.Context) int {
	users, err := r.getAllUsers(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
			fmt.Printf("%v\n", user)
		}
	return 1
}

func (r *repoService) CreateUser(ctx context.Context, input user.User) (user.User, error) {
	insert := createUserParams{
		Username: input.Username,
		UserEmail: input.UserEmail,
		UserPassword: input.UserPassword,
		FirstName: input.FirstName,
		LastName: input.LastName,
	}
	result, err := r.createUser(ctx, insert)
	if err != nil {
		return user.User{}, err
	}
	createdUser := user.User{
		UserID: result.UserID,
		Username: result.Username,
		UserEmail: result.UserEmail,
		FirstName: result.FirstName,
		LastName: result.LastName,
	}
	return createdUser, nil

}

func (r *repoService) UserByEmail(ctx context.Context, email string) (user.User, error) {
	result, err := r.getUserByEmail(ctx, email)
	if err != nil {
		return user.User{}, err
	}
	user := user.User{
		UserID: result.UserID,
		UserEmail: result.UserEmail,
		FirstName: result.FirstName,
		LastName: result.LastName,
	}
	return user, nil
}

func (r *repoService) UserByID(ctx context.Context, id string) (user.User, error) {
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return user.User{}, err
	}
	result, err := r.getUserByID(ctx, int32(idInt))
	if err != nil {
		return user.User{}, err
	}
	user := user.User{
		UserID: result.UserID,
		UserEmail: result.UserEmail,
		FirstName: result.FirstName,
		LastName: result.LastName,
	}
	return user, nil
}

func (r *repoService) UserByUsername(ctx context.Context, username string) (user.User, error) {
	result, err := r.getUserByUsername(ctx, username)
	if err != nil {
		return user.User{}, err
	}
	user := user.User{
		UserID: result.UserID,
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
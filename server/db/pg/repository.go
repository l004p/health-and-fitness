package pg

import (
	"context"
	"server/core/user"
	"server/db/pgconnect"
	"fmt"
	"log"
)

type Repository interface {
	duser.UserRepository
}

type repoService struct {
	*Queries 
	db *pgconnect.Postgres
}

func NewRepository(db *pgconnect.Postgres) Repository {
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

func (r *repoService) CreateUser(ctx context.Context, input duser.User) (int32, error) {
	insert := createUserParams{
		Username: input.Username,
		UserEmail: input.UserEmail,
		UserPassword: input.UserPassword,
		FirstName: input.FirstName,
		LastName: input.LastName,
	}
	id, err := r.createUser(ctx, insert)
	if err != nil {
		return -1, err
	}
	return id, nil

}
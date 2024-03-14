package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"server/api/graph/model"
	"strconv"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	//panic(fmt.Errorf("not implemented: CreateUser - createUser"))
	user := &model.User{
		UserID:    strconv.FormatInt(int64(1), 10),
		Username:  input.Username,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}
	return user, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	//panic(fmt.Errorf("not implemented: Users - users"))
	user1 := &model.User{
		UserID:    strconv.FormatInt(int64(1), 10),
		Username:  "testName",
		FirstName: "Justin",
		LastName:  "Foley",
	}
	user2 := &model.User{
		UserID:    strconv.FormatInt(int64(2), 10),
		Username:  "testName2",
		FirstName: "Jessica",
		LastName:  "Davies",
	}
	var users []*model.User
	users = append(users, user1, user2)
	return users, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

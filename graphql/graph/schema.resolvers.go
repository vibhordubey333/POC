package graph

//go:generate go run github.com/99designs/gqlgen --verbose

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/poc/graphql/graph/generated"
	"github.com/poc/graphql/graph/model"
)

var meetups = []*model.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A secondmeetup",
		Description: "A second description",
		UserID:      "2",
	},
}

var users = []*model.User{
	{
		ID:       "1",
		Username: "bob",
		Email:    "test@test.com",
	},
	{
		ID:       "1",
		Username: "mike",
		Email:    "test@test.com",
	},
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return meetups, nil
}

func (r *mutationResolver) User(ctx context.Context, input model.User) (*model.User, error) {
	return nil, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

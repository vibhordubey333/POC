package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/poc/graphql/graph/generated"
	"github.com/poc/graphql/graph/model"
)

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	log.Println("Request received by CreateMeetup")
	// TODO : Save in DB
	log.Println("Request successfully processed  by CreateMeetup")
	return &model.Meetup{

		Name:        input.Name,
		Description: input.Description,
	},nil
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return meetups, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type meetupResolver struct{ *Resolver }

var meetups = []*model.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		User: &model.User{
			ID:       "User_1",
			Username: "user1",
			Email:    "user1@test.com",
			Meetups:  []*model.Meetup{
				&model.Meetup{
					ID:          "1",
					Name:        "Meetup1",
					Description: "Desc_1",

				},
			},
		},
	},
	{
		ID:          "2",
		Name:        "A secondmeetup",
		Description: "A second description",
		User: &model.User{
			ID:       "",
			Username: "",
			Email:    "",
			Meetups:  nil,
		},
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

func (r *meetupResolver) User(ctx context.Context, input model.User) (*model.User, error) {
	user := new(model.User)
	log.Println("Processing User request.")
	for _, u := range users {
		if u.ID == input.ID {
			user = u
			break
		}
		if user == nil {
			return nil, errors.New("User ")
		}
	}
	log.Println("Processed User request.")
	return user, nil
}

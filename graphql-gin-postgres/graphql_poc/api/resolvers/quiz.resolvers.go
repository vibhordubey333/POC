package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vibhordubey333/graphql-gin-postgres/graphql_poc/api/generated"
	"github.com/vibhordubey333/graphql-gin-postgres/graphql_poc/api/models"
	"graphql-gin-postgres/graphql_poc/api/models"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input models.QuestionInput) (*models.Question, error) {
	panic(fmt.Errorf("not implemented"))

}

func (r *mutationResolver) CreateChoice(ctx context.Context, input *models.ChoiceInput) (*models.Choice, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Questions(ctx context.Context) ([]*models.Question, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Choices(ctx context.Context) ([]*models.Choice, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

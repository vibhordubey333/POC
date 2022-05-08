package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/vibhordubey333/POC/graphql-gin-postgres/graphql_poc/api/models"
)

func (r *mutationResolver) Quiz(ctx context.Context, luceneFilter *string, filter *models.QuestionFilterGroup, sort []*models.QuizSort, limit *models.Limit) (*models.Question, error) {
	panic(fmt.Errorf("not implemented"))
}

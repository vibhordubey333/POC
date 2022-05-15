package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	database "github.com/vibhordubey333/POC/graphql-gin-postgres/db"
	"github.com/vibhordubey333/POC/graphql-gin-postgres/graphql_poc/api/generated"
	"github.com/vibhordubey333/POC/graphql-gin-postgres/graphql_poc/api/models"
)

func (r *mutationResolver) CreateQuestion(ctx context.Context, input models.QuestionInput) (*models.Question, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	fmt.Println("input", input.QuestionText, input.PubDate)
	question := models.Question{}
	question.QuestionText = input.QuestionText
	question.PubDate = input.PubDate
	question.ID = input.ID
	db.Create(&question)
	return &question, nil
}

func (r *mutationResolver) CreateChoice(ctx context.Context, input *models.ChoiceInput) (*models.Choice, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	choice := models.Choice{}
	question := models.Question{}

	choice.QuestionID = input.QuestionID
	choice.ChoiceText = input.ChoiceText
	choice.ID = input.ChoiceID

	db.First(&question, choice.QuestionID)
	choice.Question = &question
	db.Create(&choice)
	return &choice, nil
}

func (r *queryResolver) Question(ctx context.Context) ([]*models.Question, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Find(&r.questions)
	for _, question := range r.questions {
		var choices []*models.Choice
		db.Where(&models.Choice{QuestionID: question.ID}).Find(&choices)
		question.Choices = choices
	}
	return r.questions, nil
}

func (r *queryResolver) Choices(ctx context.Context) ([]*models.Choice, error) {
	db, err := database.GetDatabase()
	if err != nil {
		log.Println("Unable to connect to database", err)
		return nil, err
	}
	defer db.Close()
	db.Find(&r.choices)
	return r.choices, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

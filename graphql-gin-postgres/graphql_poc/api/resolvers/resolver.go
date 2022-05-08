package resolvers

import "github.com/vibhordubey333/POC/graphql-gin-postgres/graphql_poc/api/models"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	questions []*models.Question
	choices	  []*models.Choice
}

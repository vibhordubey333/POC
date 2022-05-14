package main

import (
	// github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/vibhordubey333/POC/graphql-gin-postgres/graphql_poc/api/resolvers"
	"github.com/vibhordubey333/POC/graphql-gin-postgres/graphql_poc/api/generated"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8888"

func graphqlHandler() gin.HandlerFunc{
	handlerObject := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers:  &resolvers.Resolver{},
	}))
	return func(c *gin.Context){
		handlerObject.ServeHTTP(c.Writer,c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc{
	handlerObject := playground.Handler("GraphQl","/query")
	return func(c *gin.Context){
		handlerObject.ServeHTTP(c.Writer,c.Request)
	}
}

func main(){
	r := gin.Default()
	r.POST("/query",graphqlHandler())
	r.GET("/",playgroundHandler())
	r.Run(defaultPort)

}
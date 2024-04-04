package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"os"
	"vibhordubey333/POC/go-aws-serverless/pkg/handlers"
)

const (
	tableName = "go-aws-serverless"
)

var (
	dynamoClient dynamodbiface.DynamoDBAPI
)

func main() {
	region := os.Getenv("AWS_REGION")
	if region == "" {
		region = "us-east-1" // Default region if AWS_REGION is not set
	}

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	dynamoClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, tableName, dynamoClient)
	case "PUT":

	}
	return nil, nil
}

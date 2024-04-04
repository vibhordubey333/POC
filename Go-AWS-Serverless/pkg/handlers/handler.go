package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"net/http"
	"vibhordubey333/POC/go-aws-serverless/pkg/user"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMessage *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynamoDBClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	{
		email := req.QueryStringParameters["email"]
		if len(email) > 0 {
			result, err := user.FetchUser(email, tableName, dynamoDBClient)
			if err != nil {
				return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
			}
			return apiResponse(http.StatusOK, result)
		}

		result, err := user.FetchUsers(tableName, dynamoDBClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}

		return apiResponse(http.StatusOK, result)
	}
}

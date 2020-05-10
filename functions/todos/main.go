package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"math/rand"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var TodoList []Todo

func init() {
	todo1 := Todo{ID: "a", Text: "A todo not to forget", Done: false}
	todo2 := Todo{ID: "b", Text: "This is the most important", Done: false}
	todo3 := Todo{ID: "c", Text: "Please do this or else", Done: false}
	TodoList = append(TodoList, todo1, todo2, todo3)

	rand.Seed(time.Now().UnixNano())
}

func executeQuery(query string, schema graphql.Schema) (error, *graphql.Result) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if result.HasErrors() {
		return fmt.Errorf("wrong result, unexpected errors: %v", result.Errors), nil
	}
	return nil, result
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get the graphql query
	query := request.QueryStringParameters["query"]
	err, result := executeQuery(query, schema)
	statusCode := http.StatusOK
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	// Convert the output int JSON format
	jsonBytes, err := json.Marshal(result.Data)
	if err != nil {
		statusCode = http.StatusInternalServerError
	}

	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       fmt.Sprintf("%s", string(jsonBytes)),
	}, err
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

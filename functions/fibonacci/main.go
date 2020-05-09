package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bldulam1/aws-lambda-go-example/src/fib"
	"net/http"
	"strconv"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get API query
	nStr := request.QueryStringParameters["n"]
	n, err := strconv.Atoi(nStr)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "Failed to parse n",
		}, err
	}

	// Calculate fibonacci sequence
	fibN := fib.Fibonacci(n)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("Fibonacci(%s) = %d", nStr, fibN),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

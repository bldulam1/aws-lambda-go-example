package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bldulam1/aws-lambda-go-example/src/fib"
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
	factorialN := fib.Factorial(n)
	fmt.Println(n, nStr, factorialN)
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       fmt.Sprintf("Factorial(%s) = %d", nStr, factorialN),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

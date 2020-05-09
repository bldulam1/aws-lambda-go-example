package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"strconv"
)

var mem = map[int64]int64{0: 1, 1: 1, 2: 2}

func fibonacci(n int64) int64 {
	if val, ok := mem[n]; ok {
		return val
	}

	return fibonacci(n-1) + fibonacci(n)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	val := request.QueryStringParameters["val"]
	valInt, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal("Error converting into int", err)
	}

	fib := fibonacci(int64(valInt))

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf("Fib(%d)=%d", valInt, fib),
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}

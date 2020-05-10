package main

import (
	"fmt"
	"github.com/bldulam1/aws-lambda-go-example/src/fib"
)

func Sum(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(fib.Factorial(4))
}

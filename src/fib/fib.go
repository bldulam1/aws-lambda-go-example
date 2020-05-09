package fib

func Fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return n + Fibonacci(n-1)
}

func Factorial(n int) int {
	if n < 3 {
		return n
	}
	return n * Factorial(n-1)
}

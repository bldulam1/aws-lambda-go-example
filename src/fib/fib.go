package fib

var fibMem = map[int]int{
	0: 0, 1: 1, 2: 1,
}

var factMem = map[int]int{
	0: 1, 1: 1, 2: 2,
}

func Fibonacci(n int) int {
	if _, ok := fibMem[n]; !ok {
		fibMem[n] = Fibonacci(n-2) + Fibonacci(n-1)
	}
	return fibMem[n]
}

func Factorial(n int) int {
	if _, ok := factMem[n]; !ok {
		factMem[n] = n * Factorial(n-1)
	}
	return factMem[n]
}

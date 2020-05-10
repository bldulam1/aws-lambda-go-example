package fib

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	tables := []struct {
		x int
		y int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{33, 3524578},
		{39, 63245986},
		{53, 53316291173},
	}

	for _, table := range tables {
		fib := Fibonacci(table.x)
		if fib != table.y {
			t.Errorf("FAILED: Fib(%d) == %d != %d", table.x, table.y, fib)
		}
	}
}

func TestFactorial(t *testing.T) {
	tables := []struct {
		x int
		y int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{7, 5040},
		{8, 40320},
		{9, 362880},
		{10, 3628800},
	}

	for _, table := range tables {
		fib := Factorial(table.x)
		if fib != table.y {
			t.Errorf("FAILED: Fib(%d) == %d != %d", table.x, table.y, fib)
		}
	}
}

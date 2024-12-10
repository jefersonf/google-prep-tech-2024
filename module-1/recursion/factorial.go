package recursion

import (
	"fmt"
	"math"
)

func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

func Fibonacci(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func Hanoi(n int) int {
	return int(math.Pow(2, float64(n))) - 1
	// if n == 0 {
	// 	return 0
	// }
	// return 2*Hanoi(n-1) + 1
}

func HanoiMoves(n int, a, b, c rune) {
	if n >= 1 {
		HanoiMoves(n-1, a, c, b)
		fmt.Printf("move %d from %c to %c\n", n, a, c)
		HanoiMoves(n-1, b, a, c)
	}
}

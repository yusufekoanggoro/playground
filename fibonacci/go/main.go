package main

import "fmt"

func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	}

	fib := make([]int, n)
	fib[0] = 0
	fib[1] = 1

	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}

	return fib
}

func main() {
	n := 10
	result := fibonacci(n)
	fmt.Println("Deret Fibonacci:", result)
}

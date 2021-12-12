package main

import (
	"fmt"
)

var memo = make([]int, 100)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fib(n))
}

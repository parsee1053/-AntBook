package main

import (
	"fmt"
)

var n, k int
var a = make([]int, 100)

func dfs(i int, sum int) bool {
	if i == n {
		return sum == k
	}
	if dfs(i+1, sum) {
		return true
	}
	if dfs(i+1, sum+a[i]) {
		return true
	}
	return false
}

func main() {
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Scan(&k)
	if dfs(0, 0) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

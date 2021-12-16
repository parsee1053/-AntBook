package main

import (
	"fmt"
)

var n, m int
var field [][]rune

func dfs(x, y int) {
	field[x][y] = '.'
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy
			if 0 <= nx && nx < n && 0 <= ny && ny < m && field[nx][ny] == 'W' {
				dfs(nx, ny)
			}
		}
	}
}

func main() {
	fmt.Scan(&n, &m)
	field = make([][]rune, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		field[i] = []rune(s)
	}
	res := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if field[i][j] == 'W' {
				dfs(i, j)
				res++
			}
		}
	}
	fmt.Println(res)
}

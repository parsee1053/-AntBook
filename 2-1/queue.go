package main

import (
	"fmt"
)

type Queue []int

func (que *Queue) push(i int) {
	*que = append(*que, i)
}

func (que *Queue) pop() int {
	res := (*que)[0]
	*que = (*que)[1:]
	return res
}

func (que *Queue) front() int {
	return (*que)[0]
}

func main() {
	que := make(Queue, 0)
	que.push(1)
	que.push(2)
	que.push(3)
	fmt.Println(que.front())
	que.pop()
	fmt.Println(que.front())
	que.pop()
	fmt.Println(que.front())
	que.pop()
}

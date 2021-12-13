package main

import (
	"fmt"
)

type Stack []int

func (st *Stack) push(i int) {
	*st = append(*st, i)
}

func (st *Stack) pop() {
	*st = (*st)[:len(*st)-1]
}

func (st *Stack) top() int {
	return (*st)[len(*st)-1]
}

func main() {
	st := make(Stack, 0)
	st.push(1)
	st.push(2)
	st.push(3)
	fmt.Println(st.top())
	st.pop()
	fmt.Println(st.top())
	st.pop()
	fmt.Println(st.top())
	st.pop()
}

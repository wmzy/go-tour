package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		Walk(t.Left, ch)
		ch <- t.Value
		Walk(t.Right, ch)
	}
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if (<-ch1) != (<-ch2) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("T(1) is same with T(1)? %v\n", Same(tree.New(1), tree.New(1)))
	fmt.Printf("T(1) is same with T(2)? %v\n", Same(tree.New(1), tree.New(2)))
}

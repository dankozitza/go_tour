package main

import (
	"fmt"
	"code.google.com/p/go-tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	Walk(t1, c1)
	Walk(t2, c2)

	for {
		fmt.Println(<-c1, " ", <-c2)
	}
}

func main() {

	ftree := tree.New(1)
	fmt.Println(ftree.String())

	stree := tree.New(1)
	fmt.Println(stree.String())

	fmt.Println(Same(ftree, stree))
}

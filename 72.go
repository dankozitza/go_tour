package main

import (
	"fmt"
	"code.google.com/p/go-tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	for {
		select {
		case c <- 
}

// Same determines whether the trees
// t1 and t2 contain the same values.
//func Same(t1, t2 *tree.Tree) bool

func main() {

	ftree := tree.New(1)
	fmt.Println(ftree.String())

	fchan := make(chan int)

	stree := tree.New(1)
	fmt.Println(stree.String())

	schan := make(chan int)
}

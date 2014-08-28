package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var fib int
	last1 := 1
	last2 := 0
	first := true
	return func() int {
		if first {
			first = false
			return 1
		}
		fib = last1 + last2
		last2 = last1
		last1 = fib
		return fib
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}


package main

import (
   "code.google.com/p/go-tour/pic"
   "fmt"
)

func Pic(dx, dy int) [][]uint8 {
   //var graph [][]uint8

   graph := [][]uint8{{1, 2, 3}, {1, 2, 3}, {1, 2, 3},}
   fmt.Println(graph, len(graph), cap(graph))
   //for i := 0; i < 10; i++ {
   //   var pix uint8 = 1
   //   for o := 0; o < 10; o++ {
   //      graph[i][o] = pix
   //      pix++
   //   }
   //}
   return graph
}

func main() {
    pic.Show(Pic)
}

package main

import (
   "code.google.com/p/go-tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

   graph := make([][]uint8, dy)

   for y := range graph {
      graph[y] = make([]uint8, dx)
      for x := range graph[y] {
         graph[y][x] = uint8(x * y)
      }
   }

   return graph
}

func main() {
    pic.Show(Pic)
}

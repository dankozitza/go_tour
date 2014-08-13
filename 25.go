package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
   var z float64 = 1
   var last float64 = 0
   var loops int = 1
   for {
      var diff float64 = last - z
      if diff < 0 {
         diff = diff * -1;
      }

      if diff > .01 {
         last = z
         z = z - (((x * x) - x) / (2 * x))
         loops = loops + 1
         fmt.Println(z)
      } else {
         break
      }
   }

   fmt.Println("loops: ")
   fmt.Println(loops)

   return z
}

func main() {
    fmt.Println(Sqrt(4))
    fmt.Println("\n")
}

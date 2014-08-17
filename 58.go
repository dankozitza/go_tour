package main

import (
   "fmt"
)

func Sqrt(x float64) (float64, error) {
   var z float64 = 1
   var last float64 = 2
   var cnt int64 = 0

   if (x < 0) {
      var err ErrNegativeSqrt
      return 0, err
   }

   for {
      var diff float64 = last - z
      if diff < 0 {
         diff = diff * -1;
      }

      if diff > .01 {
         last = z
         z = z - (((z * z) - x) / (2 * z))
         cnt = cnt + 1
         fmt.Println("cnt: [", cnt, "] number: [", z, "] diff: [", diff, "]")
      } else {
         break
      }
   }

   return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
   return "cannot get square root of negative number!"
}

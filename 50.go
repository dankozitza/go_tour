package main

import (
    "fmt"
	 "math/cmplx"
)

func Cbrt(x complex128) complex128 {
   var z complex128 = 1
   var last complex128 = 2
   var cnt int = 0

   for {
      last = z
      z = z - (((cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2))))
      cnt += 1
      
      var diff complex128 = last - z
		fmt.Println("cnt: [", cnt, "] number: [", z, "] diff: [", diff, "]")

      if cmplx.Abs(diff) < .01 {
			break
		}
   }

   return z
}

func main() {
    fmt.Println("approximation: ", Cbrt(2))
}

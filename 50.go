package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func Cbrt(x complex128) complex128 {
	var z complex128 = 1
	var last complex128 = 2
	var cnt int = 0

	for {
		last = z
		z = z - ((cmplx.Pow(z, 3) - x) / (3 * cmplx.Pow(z, 2)))
		cnt += 1

		var diff complex128 = last - z
		var r, t float64 = cmplx.Polar(diff)

		fmt.Println("cnt: [", cnt, "] z: [", z, "] r: [", r, "] t: [", t, "]")

		if math.Abs(r) < .00000001 && math.Abs(t) < .00000001 {
			break
		}
	}

	return z
}

func main() {
	var in complex128 = 2
	fmt.Println("finding cube root for complex number: ", in)
	fmt.Println("approximation: ", Cbrt(in))
}

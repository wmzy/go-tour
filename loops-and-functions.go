package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, n int) float64 {
	z := 1.0
	for ; n > 0; n-- {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2, 2), math.Sqrt(2))
	fmt.Println(Sqrt(2, 20), math.Sqrt(2))
}

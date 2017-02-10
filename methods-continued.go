package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (v MyFloat) Abs() float64 {
	if v < 0 {
		return float64(-v)
	}
	return float64(v)
}

func main() {
	v := MyFloat(-math.Sqrt2)
	fmt.Println(v.Abs())
	v = 1.0
	fmt.Println(v.Abs())
}

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for mid := (z*z - x) / (2 * z); math.Abs(mid) > 0.0000001; mid = (z*z - x) / (2 * z) {
		z -= mid
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}

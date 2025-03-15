package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%.9f", sqrt(2))
}

func sqrt(x float64) float64 {
	acc := 0.00001
	l := 0.0
	r := x
	ans := (l + r) / 2
	for math.Abs(ans*ans-x) > acc {
		// update ans
		if ans*ans < x {
			l = ans
		} else {
			r = ans
		}
		ans = (l + r) / 2
	}
	return ans
}

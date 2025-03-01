package main

import (
	"fmt"
	"math/rand"
)

func inCircle(x, y float64) bool {
	return x*x+y*y <= 1.0
}

func main() {
	n := 100_000

	m := 0.0
	for i := 0; i < n; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if inCircle(x, y) {
			m += 1.0
		}
	}

	var ans float64 = 4.0 * m / float64(n)
	fmt.Printf("pi: %.6f\n", ans)
}

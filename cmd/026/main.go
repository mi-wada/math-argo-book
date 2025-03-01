package main

import (
	"fmt"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	ans := 0.0
	for i := 0; i < n; i++ {
		ans += float64(n) / (float64(n) - float64(i))
	}
	fmt.Printf("%.6f\n", ans)
}

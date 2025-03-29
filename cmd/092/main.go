package main

import (
	"fmt"
	"math"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	// Nの約数をO(sqrt(N))で列挙。その約数をaとして、b=N/aとする。その時、a+bが最小となるようなa,bの組み合わせを求める。
	ans := math.MaxInt64
	for a := 1; a <= int(math.Sqrt(float64(n))); a++ {
		if n%a != 0 {
			continue
		}
		b := n / a
		ans = min(ans, a+b)
	}
	fmt.Println(ans * 2)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

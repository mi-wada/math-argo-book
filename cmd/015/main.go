package main

import (
	"fmt"
)

func main() {
	// A B
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(gcd(a, b))
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

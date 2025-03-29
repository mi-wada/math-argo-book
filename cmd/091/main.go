package main

import (
	"fmt"
)

func main() {
	// N X
	var n, x int
	fmt.Scan(&n, &x)

	var ans int
	for a := 1; a <= n; a++ {
		for b := a + 1; b <= n; b++ {
			for c := b + 1; c <= n; c++ {
				if a+b+c == x {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}

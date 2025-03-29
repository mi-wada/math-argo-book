package main

import (
	"fmt"
)

func main() {
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	var ok bool
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			for c := 1; c <= n; c++ {
				d := x - a - b - c
				if d <= 0 || d > n {
					continue
				}
				if a*b*c*d == y {
					ok = true
					break
				}
			}
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import (
	"fmt"
)

func main() {
	// a b c d
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	ans := maxInt(
		a*c,
		a*d,
		b*c,
		b*d,
	)
	fmt.Println(ans)
}

func maxInt(x int, y ...int) int {
	for _, v := range y {
		if v > x {
			x = v
		}
	}
	return x
}

package main

import (
	"fmt"
)

func main() {
	// A B
	var a, b int
	fmt.Scan(&a, &b)

	var ans int
	for i := 1; i <= b; i++ {
		if floor(b, i)-ceil(a, i) >= 1 {
			ans = i
		}
	}
	fmt.Println(ans)
}

func floor(x, y int) int {
	return x / y
}

func ceil(x, y int) int {
	if x%y == 0 {
		return x / y
	} else {
		return x/y + 1
	}
}

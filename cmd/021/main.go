package main

import (
	"fmt"
)

func main() {
	// n r
	var n, r int
	fmt.Scan(&n, &r)

	ans := 1
	for i := 0; i < r; i++ {
		ans *= (n - i)

	}
	rr := 1
	for i := 0; i < r; i++ {
		rr *= (r - i)
	}

	ans /= rr
	fmt.Println(ans)
}

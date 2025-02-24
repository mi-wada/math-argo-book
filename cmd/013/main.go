package main

import (
	"fmt"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
			fmt.Println(n / i)
		}
	}
}

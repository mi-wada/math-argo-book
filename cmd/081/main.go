package main

import (
	"fmt"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	cnt10000 := n / 10000
	n -= 10000 * cnt10000
	cnt5000 := n / 5000
	n -= 5000 * cnt5000
	cnt1000 := n / 1000
	n -= 1000 * cnt1000

	fmt.Println(cnt1000 + cnt5000 + cnt10000)
}

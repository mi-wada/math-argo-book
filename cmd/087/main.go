package main

import (
	"fmt"
)

const mod = 1000000007

func main() {
	var n int
	fmt.Scan(&n)

	sum := (n * (n + 1) / 2) % mod
	ans := sum * sum % mod
	fmt.Println(ans)
}

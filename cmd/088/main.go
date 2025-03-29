package main

import (
	"fmt"
)

const mod = 998244353

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	sum := func(n int) int { return n * (n + 1) / 2 }
	ans := (((sum(a) % mod) * (sum(b) % mod)) % mod * (sum(c) % mod)) % mod
	fmt.Println(ans)
}

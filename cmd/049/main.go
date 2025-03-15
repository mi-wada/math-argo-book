package main

import (
	"fmt"
)

const mod = 1000000007

func main() {
	// n
	var n int
	fmt.Scan(&n)

	fib := make([]int, n+1)
	fib[1] = 1
	fib[2] = 1
	for i := 3; i <= n; i++ {
		fib[i] = (fib[i-1]%mod + fib[i-2]%mod) % mod
	}

	fmt.Println(fib[n])
}

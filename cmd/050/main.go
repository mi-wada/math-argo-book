package main

import (
	"fmt"
)

const mod = 1000000007

func main() {
	// a b
	var a, b int
	fmt.Scan(&a, &b)

	fmt.Println(powpow(a, b))
}

// powpow caluculate a^b % mod
func powpow(a int, b int) int {
	ret := 1
	i := a
	for b > 0 {
		if b&1 == 1 {
			ret = (ret * i) % mod
		}
		b = b >> 1
		i = (i * i) % mod
	}
	return ret
}

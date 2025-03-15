package main

import (
	"fmt"
)

const mod = 1000000007

func main() {
	// X Y
	var x, y int
	fmt.Scan(&x, &y)

	// ans: (x + y)! / (x! * y!)
	// (x + y)!
	bunshi := 1
	for i := 1; i <= (x + y); i++ {
		bunshi = (bunshi * i) % mod
	}
	// x!
	bunboX := 1
	for i := 1; i <= x; i++ {
		bunboX = (bunboX * i) % mod
	}
	// y!
	bunboY := 1
	for i := 1; i <= y; i++ {
		bunboY = (bunboY * i) % mod
	}
	bunbo := (bunboX * bunboY) % mod

	// フェルマーの小定理より
	// (bunshi / bunbo)%mod == bunshi * powpow(bunbo, m-2)%mod
	ans := bunshi * powpow(bunbo, mod-2) % mod
	fmt.Println(ans)
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

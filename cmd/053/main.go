package main

import (
	"fmt"
)

const mod = 1000000007

func main() {
	// N
	var n int
	fmt.Scan(&n)

	// 等比数列の和の公式
	// https://manabitimes.jp/math/948
	//
	// フェルマーの小定理より
	// (a/b)%mod == a*powpow(b, mod-2)%mod
	ans := (powpow(4, n+1) - 1) * powpow(3, mod-2) % mod
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

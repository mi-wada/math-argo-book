package main

import (
	"fmt"
)

func main() {
	// A B
	var a, b int64
	fmt.Scan(&a, &b)

	// 0 <= a, b <= 10^18
	// aとbの最小公倍数を求める
	var ans int64

	ans = lcm(a, b)
	if ans == -1 {
		fmt.Println("Large")
		return
	}
	fmt.Println(ans)
}

// 最大公約数を求める
func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

// 最小公倍数を求める
func lcm(a, b int64) int64 {
	// オーバーフロー対策: a * b / gcd(a, b) の代わりに
	// a / gcd(a, b) * b を計算する
	g := gcd(a, b)
	// オーバーフロー対策
	a /= g
	// a*bがint64の最大値を超える場合-1を返す
	if a > 0 && b > 1000000000000000000/a {
		return -1
	}
	return a * b
}

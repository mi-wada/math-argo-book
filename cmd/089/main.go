package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int64
	fmt.Scan(&a, &b, &c)

	// 式変形により a < c^bの判定問題となる。
	// c^bを繰り返しに情報で求めて、もしそれが10^18を超えるならばあかんと判定し、そうでない場合は単純に数値の比較をする。
	if a < fastpow(c, b) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

// もし結果が10^18を超えるならば10^18 + 1を返す。
func fastpow(a, b int64) int64 {
	var MAX int64 = math.MaxInt64
	var ret int64 = 1
	for b > 0 {
		if b&1 == 1 {
			if ret > MAX/a {
				return MAX
			}
			ret = ret * a
		}
		b = b >> 1
		if a > MAX/a {
			if b > 0 {
				return MAX
			}
			break
		}
		a = a * a
	}
	return ret
}

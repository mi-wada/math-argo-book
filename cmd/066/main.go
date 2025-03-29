package main

import (
	"fmt"
)

func main() {
	// N K
	var n, k int
	fmt.Scan(&n, &k)

	// r: 赤色のカードに書かれた整数
	// w: 白色のカードに書かれた整数
	// b: 青色のカードに書かれた整数
	//
	// N^3通りありうる。全探索は無理。
	//
	// 以下の条件のいずれかを満たすものをカウント。
	// * |r - w| >= k
	// * |r - b| >= k
	// * |w - b| >= k
	//
	// 余事象は以下の条件を全て満たすものである。
	// * |r - w| < k
	// * |r - b| < k
	// * |w - b| < k
	//
	// N^3から余事象のカウントを引けば答えになる。
	// 問題は余事象のカウントは制限時間内で算出可能かどうか。
	//
	// |r - w| < kとは以下と同義である。
	// -k < r-w < k
	// -k < r-w && r-w < k
	// w < r+k && r-k < w
	//
	// |w - b| < kとは以下と同義である。
	// -k < w-b < k
	// -k < w-b && w-b < k
	// b < w+k && w-k < b

	var cnt int
	for r := 1; r <= n; r++ {
		for w := maxInt(1, r-k); w <= minInt(n, r+k); w++ {
			for b := maxInt(1, w-k); b <= minInt(n, w+k); b++ {
				if absInt(r-w) < k &&
					absInt(r-b) < k &&
					absInt(w-b) < k {
					cnt++
				}
			}
		}
	}
	ans := n*n*n - cnt
	fmt.Println(ans)
}

func absInt(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

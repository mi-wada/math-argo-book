package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	// a1, a2, ..., an
	// WARN: linesをひとつのstringにまとめて取得する方法だと実行時エラーになる
	a := make([]uint64, n)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.ParseUint(s, 10, 64)
		a[i] = v
	}

	ans := a[0]
	for _, v := range a {
		ans = gcd(v, ans)
	}
	fmt.Println(ans)
}

func gcd(a, b uint64) uint64 {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

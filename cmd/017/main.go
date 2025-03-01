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
	a := make([]int, n)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		a[i] = v
	}

	ans := a[0]
	for _, v := range a {
		ans = lcm(v, ans)
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	g := gcd(a, b)
	// WARN: a * b / gだとオーバーフローする
	return (a / g) * b
}

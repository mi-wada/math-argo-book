package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// t
	// n
	var t, n int
	fmt.Scan(&t, &n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	// l1, r1
	// ...
	// ln, rn
	lr := make([]struct {
		l int
		r int
	}, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		lr[i].l = v

		scanner.Scan()
		s = scanner.Text()
		v, _ = strconv.Atoi(s)
		lr[i].r = v
	}

	// cnt[i]: 時刻x i <= x < i+1に働いている従業員数
	// 0 <= i <= t
	cnt := make([]int, t+1)
	for i := 1; i <= n; i++ {
		cnt[lr[i].l] += 1
		cnt[lr[i].r] -= 1
	}
	for i := 1; i <= t; i++ {
		cnt[i] += cnt[i-1]
	}
	for i := 0; i < t; i++ {
		fmt.Println(cnt[i])
	}
}

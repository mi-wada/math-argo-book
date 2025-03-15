package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	scanner := bufio.NewScanner(os.Stdin)
	// a1, a2, ..., an
	a := make([]int, n+1)
	scanner.Split(bufio.ScanWords)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		a[i] = v
	}
	// l1, r1
	// ...
	// lq, rq
	lr := make([]struct {
		l int
		r int
	}, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		lr[i].l = v

		scanner.Scan()
		s = scanner.Text()
		v, _ = strconv.Atoi(s)
		lr[i].r = v
	}

	// visitedCnt[i]: i日目までに訪れた人数
	// 0 <= i <= n
	visitedCnt := make([]int, n+1)
	for i := 1; i <= n; i++ {
		visitedCnt[i] = visitedCnt[i-1] + a[i]
	}

	for i := 0; i < q; i++ {
		fmt.Println(visitedCnt[lr[i].r] - visitedCnt[lr[i].l-1])
	}
}

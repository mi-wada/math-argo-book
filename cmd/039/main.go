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
	scanner.Split(bufio.ScanWords)
	// l1 r1 x1
	// ...
	// lq rq xq
	lrx := make([]struct {
		l int
		r int
		x int
	}, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		lrx[i].l = v

		scanner.Scan()
		s = scanner.Text()
		v, _ = strconv.Atoi(s)
		lrx[i].r = v

		scanner.Scan()
		s = scanner.Text()
		v, _ = strconv.Atoi(s)
		lrx[i].x = v
	}

	snow := make([]int, n+2)
	for i := 0; i < q; i++ {
		snow[lrx[i].l] += lrx[i].x
		snow[lrx[i].r+1] -= lrx[i].x
	}
	for i := 1; i <= n; i++ {
		snow[i] += snow[i-1]
	}

	for i := 1; i < n; i++ {
		var ans string
		switch {
		case snow[i] > snow[i+1]:
			ans = ">"
		case snow[i] == snow[i+1]:
			ans = "="
		default:
			ans = "<"
		}
		fmt.Print(ans)
	}
	fmt.Println()
}

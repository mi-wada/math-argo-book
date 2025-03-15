package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// N W
	var n, w int
	fmt.Scan(&n, &w)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	// w1, v1
	// w2, v2
	// ...
	// wn, vn
	wv := make([]struct{ w, v int }, n+1)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		wv[i].w = v

		scanner.Scan()
		s = scanner.Text()
		v, _ = strconv.Atoi(s)
		wv[i].v = v
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for ww := 1; ww <= w; ww++ {
			if wv[i].w <= ww {
				dp[i][ww] = maxInt(
					dp[i-1][ww],
					dp[i-1][ww-wv[i].w]+wv[i].v,
				)
			} else {
				dp[i][ww] = dp[i-1][ww]
			}
		}
	}

	fmt.Println(dp[n][w])
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

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
	a := make([]int, n)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		a[i] = v
	}

	dp := make([]int, n+1)
	dp[1] = a[0]
	for i := 2; i <= n; i++ {
		dp[i] = maxInt(dp[i-1], dp[i-2]+a[i-1])
	}
	fmt.Println(dp[n])
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

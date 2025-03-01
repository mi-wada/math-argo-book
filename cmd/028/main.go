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

	dp := make([]int, len(a))
	dp[0] = 0
	dp[1] = absInt(a[1] - a[0])
	for i := 2; i < n; i++ {
		dp[i] = minInt(
			dp[i-1]+absInt(a[i]-a[i-1]),
			dp[i-2]+absInt(a[i]-a[i-2]),
		)
	}

	fmt.Println(dp[len(dp)-1])
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func absInt(a int) int {
	if a >= 0 {
		return a
	} else {
		return a * -1
	}
}

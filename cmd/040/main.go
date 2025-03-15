package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)
	scanner := bufio.NewScanner(os.Stdin)
	// a1, a2, ..., an-1
	a := make([]int, n)
	scanner.Split(bufio.ScanWords)
	for i := 1; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		a[i] = v
	}
	// m
	// b1
	// ...
	// bm
	var m int
	scanner.Scan()
	s := scanner.Text()
	v, _ := strconv.Atoi(s)
	m = v
	b := make([]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		b[i] = v
	}

	// c[i]: 駅1から駅iにいくまでに必要な距離
	// 1 <= i <= n
	c := make([]int, n+1)
	for i := 2; i <= n; i++ {
		c[i] = c[i-1] + a[i-1]
	}

	var ans int
	for i := 1; i < m; i++ {
		// 駅b[i-1] -> b[i]
		ans += absInt(c[b[i-1]] - c[b[i]])
	}
	fmt.Println(ans)
}

func absInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

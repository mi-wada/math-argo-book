package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// n m
	var n, m int
	fmt.Scan(&n, &m)

	// c[i]: ノードiへの、自分自身より小さい頂点番号からの入次数
	c := make([]int, n)

	// a1 b1
	// ...
	// an bn
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		a--
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		b--

		if a > b {
			c[a]++
		} else {
			c[b]++
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		if c[i] == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}

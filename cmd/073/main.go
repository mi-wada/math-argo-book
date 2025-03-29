package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func main() {
	// N
	var n int
	fmt.Scan(&n)

	// A1 ... An
	a := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	var ans int
	for i := 0; i < n; i++ {
		ans += a[i] * (2*i - n + 1)
	}
	fmt.Println(ans)
}

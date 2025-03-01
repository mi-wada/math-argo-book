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
	scanner.Split(bufio.ScanWords)
	// A1, A2, ..., An
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		a[i] = v
	}
	// B1, B2, ..., Bn
	b := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		b[i] = v
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		ans += float64(a[i]) / 3.0
		ans += float64(b[i]) * 2.0 / 3.0
	}
	fmt.Printf("%.6f\n", ans)
}

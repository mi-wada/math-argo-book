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
	// B1, B2, ..., Bn
	b := make([]int, n)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		b[i] = v
	}
	// R1, R2, ..., Rn
	r := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		r[i] = v
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		ans += float64(b[i])
		ans += float64(r[i])
	}
	ans /= float64(n)
	fmt.Printf("%.6f\n", ans)
}

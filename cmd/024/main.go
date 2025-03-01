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
	// P1, Q1
	// P2, Q2
	// ...
	// Pn, Qn
	pq := make([]struct{ p, q int }, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		pq[i].p = v

		scanner.Scan()
		s = scanner.Text()
		v, _ = strconv.Atoi(s)
		pq[i].q = v
	}

	ans := 0.0
	for i := 0; i < n; i++ {
		ans += float64(pq[i].q) / float64(pq[i].p)
	}
	fmt.Printf("%.6f\n", ans)
}

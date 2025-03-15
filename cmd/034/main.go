package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type cord struct {
	x, y float64
}

func main() {
	var n int
	fmt.Scan(&n)
	s := make([]cord, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s[i].x, _ = strconv.ParseFloat(scanner.Text(), 64)

		scanner.Scan()
		s[i].y, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	ans := math.MaxFloat64
	for i := 0; i < n; i++ {
		for j := 0; j < n && j != i; j++ {
			dist := math.Sqrt(math.Pow(s[i].x-s[j].x, 2) + math.Pow(s[i].y-s[j].y, 2))
			if dist < ans {
				ans = dist
			}
		}
	}

	fmt.Printf("%.9f\n", ans)
}

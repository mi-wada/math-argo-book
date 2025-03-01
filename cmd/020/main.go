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

	m := make([]map[int]int, n)
	for i := range m {
		m[i] = make(map[int]int)
	}
	m[n-1][a[n-1]]++
	for i := n - 2; i >= 0; i-- {
		for k, v := range m[i+1] {
			m[i][k] = v
		}
		m[i][a[i]]++
	}

	ans := 0
	for i1 := 0; i1 < n-4; i1++ {
		for i2 := i1 + 1; i2 < n-3; i2++ {
			for i3 := i2 + 1; i3 < n-2; i3++ {
				for i4 := i3 + 1; i4 < n-1; i4++ {
					sum := a[i1] + a[i2] + a[i3] + a[i4]
					ans += m[i4+1][1000-sum]
				}
			}
		}
	}
	fmt.Println(ans)
}

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

	m := make(map[int]int)
	for _, v := range a {
		m[v]++
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		m[a[i]]--
		ans += m[100000-a[i]]
	}
	fmt.Println(ans)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// N X
	var n, x int
	fmt.Scan(&n, &x)
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

	sort.Ints(a)

	if binarySearch(a, x) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func binarySearch(s []int, x int) bool {
	l, r := 0, len(s)-1
	for 1 < r-l {
		m := (l + r) / 2
		if s[m] < x {
			l = m
		} else {
			r = m
		}
	}
	return s[l] == x || s[r] == x
}

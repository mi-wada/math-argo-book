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

	mergeSort(a)

	for i := 0; i < len(a); i++ {
		if i == len(a)-1 {
			fmt.Printf("%d\n", a[i])
		} else {
			fmt.Printf("%d ", a[i])
		}
	}
}

func mergeSort(s []int) {
	if len(s) <= 1 {
		return
	}

	l, r := split(s)
	mergeSort(l)
	mergeSort(r)

	li, ri := 0, 0
	for i := 0; i < len(s); i++ {
		switch {
		case li == len(l):
			s[i] = r[ri]
			ri++
		case ri == len(r):
			s[i] = l[li]
			li++
		default:
			if l[li] < r[ri] {
				s[i] = l[li]
				li++
			} else {
				s[i] = r[ri]
				ri++
			}
		}
	}
}

func split(s []int) ([]int, []int) {
	lLen := len(s) / 2
	rLen := len(s) - lLen
	l, r := make([]int, lLen), make([]int, rLen)
	for i := 0; i < lLen; i++ {
		l[i] = s[i]
	}
	for i := 0; i < rLen; i++ {
		r[i] = s[lLen+i]
	}
	return l, r
}

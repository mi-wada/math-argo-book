package main

import (
	"cmp"
	"fmt"
)

func main() {
	s := []int{5, 1, 2, 4, 3}
	mergeSort(s)
	fmt.Println(s)
}

func selectionSort[S ~[]E, E cmp.Ordered](s S) {
	for i := 0; i < len(s); i++ {
		minI := i
		for j := i; j < len(s); j++ {
			if s[j] < s[minI] {
				minI = j
			}
		}
		s[i], s[minI] = s[minI], s[i]
	}
}

func mergeSort[S ~[]E, E cmp.Ordered](s S) {
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

func split[S ~[]E, E cmp.Ordered](s S) (S, S) {
	lLen := len(s) / 2
	rLen := len(s) - lLen
	l, r := make(S, lLen), make(S, rLen)
	for i := 0; i < lLen; i++ {
		l[i] = s[i]
	}
	for i := 0; i < rLen; i++ {
		r[i] = s[lLen+i]
	}
	return l, r
}

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	ans := factorization(n)

	ansS := make([]string, len(ans))
	for i := range ans {
		ansS[i] = strconv.Itoa(ans[i])
	}
	fmt.Println(strings.Join(ansS, " "))
}

func factorization(x int) []int {
	s := make([]int, 0)
	if x <= 1 {
		return s
	}
	divisor := smallestDivisor(x)
	s = append(s, divisor)
	s = append(s, factorization(x/divisor)...)
	return s
}

// smallestDivisor returns the smallest divisor of x that is 2 or greater.
func smallestDivisor(x int) int {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return i
		}
	}
	return x
}

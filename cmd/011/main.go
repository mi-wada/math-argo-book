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

	ans := make([]string, 0)
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	fmt.Println(strings.Join(ans, " "))
}

func isPrime(x int) bool {
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

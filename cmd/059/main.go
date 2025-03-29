package main

import (
	"fmt"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	m := map[int]int{
		1: 2,
		2: 4,
		3: 8,
		4: 6,
	}
	ans := m[(n-1)%4+1]
	fmt.Println(ans)
}

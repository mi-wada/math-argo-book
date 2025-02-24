package main

import "fmt"

func main() {
	// N
	var n int
	fmt.Scan(&n)

	ans := 1
	for i := 1; i <= n; i++ {
		ans *= i
	}
	fmt.Println(ans)
}

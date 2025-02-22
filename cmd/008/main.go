package main

import "fmt"

func main() {
	// N S
	var n, s int
	fmt.Scan(&n, &s)

	ans := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i+j <= s {
				ans += 1
			}
		}
	}
	fmt.Println(ans)
}

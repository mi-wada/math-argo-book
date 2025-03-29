package main

import (
	"fmt"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	if n%2 == 0 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

package main

import (
	"fmt"
)

const mod = 1000000007

func main() {
	// N X Y
	var n, x, y int
	fmt.Scan(&n, &x, &y)

	var ok bool
	for b := 0; b <= n; b++ {
		u := y + b
		r := x - (u + b - n)
		if r%2 == 0 {
			r /= 2
		} else {
			continue
		}
		l := n - (u + r + b)
		if u >= 0 && r >= 0 && b >= 0 && l >= 0 {
			ok = true
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

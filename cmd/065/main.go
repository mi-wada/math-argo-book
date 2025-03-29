package main

import (
	"fmt"
)

func main() {
	// H W
	var h, w int
	fmt.Scan(&h, &w)

	if h == 1 || w == 1 {
		fmt.Println(1)
		return
	}

	ceilW, floorW := intCeil(w), intFloor(w)
	ceilH, floorH := intCeil(h), intFloor(h)

	ans := ceilW*ceilH + floorW*floorH
	fmt.Println(ans)
}

func intCeil(n int) int {
	if n%2 == 0 {
		return n / 2
	} else {
		return n/2 + 1
	}
}

func intFloor(n int) int {
	return n / 2
}

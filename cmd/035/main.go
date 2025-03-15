package main

import (
	"fmt"
)

type circle struct {
	x, y int64
	r    int64
}

func main() {
	var c1 circle
	fmt.Scan(&c1.x, &c1.y, &c1.r)
	var c2 circle
	fmt.Scan(&c2.x, &c2.y, &c2.r)
	if c1.r < c2.r {
		c1, c2 = c2, c1
	}

	d := (c1.x-c2.x)*(c1.x-c2.x) + (c1.y-c2.y)*(c1.y-c2.y)

	var ans int
	switch {
	case d < c1.r*c1.r:
		ans = 1
	case d == (c1.r-c2.r)*(c1.r-c2.r):
		ans = 2
	case d < (c1.r+c2.r)*(c1.r+c2.r):
		ans = 3
	case d == (c1.r+c2.r)*(c1.r+c2.r):
		ans = 4
	case (c1.r+c2.r)*(c1.r+c2.r) < d:
		ans = 5
	}
	fmt.Println(ans)
}

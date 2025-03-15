package main

import (
	"fmt"
	"math"
)

type cord struct {
	x, y float64
}

func main() {
	var a cord
	fmt.Scan(&a.x, &a.y)
	var b cord
	fmt.Scan(&b.x, &b.y)
	var c cord
	fmt.Scan(&c.x, &c.y)

	abcOver90 := (a.x-b.x)*(c.x-b.x)+(a.y-b.y)*(c.y-b.y) < 0
	acbOver90 := (b.x-c.x)*(a.x-c.x)+(b.y-c.y)*(a.y-c.y) < 0
	var ans float64
	if abcOver90 {
		ans = distance(a, b)
	} else if acbOver90 {
		ans = distance(a, c)
	} else {
		ba := cord{x: a.x - b.x, y: a.y - b.y}
		bc := cord{x: c.x - b.x, y: c.y - b.y}
		babc := int(math.Abs(float64(ba.x*bc.y - ba.y*bc.x)))
		ans = float64(babc) / distance(b, c)
	}
	fmt.Printf("%.6f\n", ans)
}

func distance(a, b cord) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	return math.Sqrt(dx*dx + dy*dy)
}

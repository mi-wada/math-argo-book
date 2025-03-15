package main

import (
	"fmt"
)

type cord struct {
	x, y float64
}

func main() {
	var c1, c2, c3, c4 cord
	fmt.Scan(&c1.x, &c1.y)
	fmt.Scan(&c2.x, &c2.y)
	fmt.Scan(&c3.x, &c3.y)
	fmt.Scan(&c4.x, &c4.y)

	// a1 := (c2.y - c1.y) / (c2.x - c1.x)
	// b1 := c1.y - a1*c1.x
	// f1 := func(x float64) float64 {
	// 	return a1*x + b1
	// }

	// a2 := (c4.y - c3.y) / (c4.x - c3.x)
	// b2 := c3.y - a2*c3.x
	// f2 := func(x float64) float64 {
	// 	return a2*x + b2
	// }

	// // 並行ならNo。ただしこの判定は誤差を含むため多分間違っている。
	// if a1 == a2 {
	// 	fmt.Println("No")
	// 	return
	// }

	// var cross cord
	// cross.x = -(b1 - b2) / (a1 - a2)
	// cross.y = f1(cross.x)
}

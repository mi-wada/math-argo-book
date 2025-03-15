package main

import (
	"fmt"
	"math"
)

type cord struct {
	x, y int
}

func main() {
	// x1 y1
	var a, b, h, m int
	fmt.Scan(&a, &b, &h, &m)
	var thetaA float64 = 360*float64(h)/12 + 30*float64(m)/60
	var thetaB float64 = 360 * float64(m) / 60
	var thetaAB float64 = math.Abs(thetaA - thetaB)
	var thetaABRad float64 = 2 * math.Pi * thetaAB / 360
	var ans float64 = float64(a)*float64(a) + float64(b)*float64(b) - 2*float64(a)*float64(b)*math.Cos(thetaABRad)
	ans = math.Sqrt(ans)
	fmt.Printf("%.9f\n", ans)
}

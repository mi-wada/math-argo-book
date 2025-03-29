package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type rect struct {
	l, r, b, t int
}

func (r rect) Contains(x, y int) bool {
	return r.l <= x && x <= r.r && r.b <= y && y <= r.t
}

func (r rect) S() int {
	return (r.t - r.b) * (r.r - r.l)
}

func main() {
	// N K
	var n, k int
	fmt.Scan(&n, &k)

	xy := make([]struct{ x, y int }, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		xy[i].x, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		xy[i].y, _ = strconv.Atoi(scanner.Text())
	}

	ans := math.MaxInt
	for li := 0; li < n; li++ {
		for ri := 0; ri < n; ri++ {
			for bi := 0; bi < n; bi++ {
				for ti := 0; ti < n; ti++ {
					r := rect{
						l: xy[li].x,
						r: xy[ri].x,
						b: xy[bi].y,
						t: xy[ti].y,
					}
					var cnt int
					for i := 0; i < n; i++ {
						if r.Contains(xy[i].x, xy[i].y) {
							cnt++
						}
					}
					if cnt == k {
						ans = minInt(ans, r.S())
					}
				}
			}
		}
	}
	fmt.Println(ans)
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

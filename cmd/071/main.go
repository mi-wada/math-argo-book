package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	abc := make([]struct{ a, b, c float64 }, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		abc[i].a, _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		abc[i].b, _ = strconv.ParseFloat(scanner.Text(), 64)
		scanner.Scan()
		abc[i].c, _ = strconv.ParseFloat(scanner.Text(), 64)
	}

	ans := math.Inf(-1)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x := (abc[i].c*abc[j].b - abc[j].c*abc[i].b) / (abc[i].a*abc[j].b - abc[j].a*abc[i].b)
			y := (abc[i].c*abc[j].a - abc[j].c*abc[i].a) / (abc[i].b*abc[j].a - abc[j].b*abc[i].a)

			ok := true
			for k := 0; k < n; k++ {
				if abc[k].a*x+abc[k].b*y > abc[k].c {
					ok = false
					break
				}
			}
			if ok {
				ans = math.Max(ans, x+y)
			}
		}
	}

	fmt.Println(ans)
}

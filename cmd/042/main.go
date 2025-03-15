package main

import (
	"fmt"
)

func main() {
	// n
	var n int64
	fmt.Scan(&n)

	// divisors[i]: 整数iの約数の数
	divisors := make([]int64, n+1)
	for i := 0; i < int(n)+1; i++ {
		divisors[i] = 1
	}
	divisors[0] = 1
	divisors[1] = 1
	for i := 2; i <= int(n); i++ {
		if divisors[i] != 1 {
			continue
		}
		for j := i; j <= int(n); j += i {
			k := 1
			copiedJ := j
			for copiedJ%i == 0 {
				copiedJ /= i
				k++
			}
			divisors[j] *= int64(k)
		}
	}

	// for i, v := range divisors {
	// 	fmt.Printf("%d: %d\n", i, v)
	// }

	var ans int64
	for i := 1; i <= int(n); i++ {
		ans += int64(i) * divisors[i]
	}
	fmt.Println(ans)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// N K
	var n, k int
	fmt.Scan(&n, &k)

	v := make([]int, k)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < k; i++ {
		scanner.Scan()
		v[i], _ = strconv.Atoi(scanner.Text())
	}

	var ans int
	for i := 1; i < 1<<k; i++ {
		var l int
		for bit := 0; bit < k; bit++ {
			if (i>>bit)&1 == 1 {
				if l == 0 {
					l = v[bit]
				} else {
					l = lcm(l, v[bit])
				}
			}
		}
		if oneCount(i)%2 == 0 {
			ans -= n / l
		} else {
			ans += n / l
		}
	}
	fmt.Println(ans)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return a / gcd(a, b) * b
}

func oneCount(n int) int {
	if n == 0 {
		return 0
	}

	if n&1 == 1 {
		return 1 + oneCount(n>>1)
	} else {
		return oneCount(n >> 1)
	}
}

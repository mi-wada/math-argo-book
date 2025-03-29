package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const mod = 1000000007

func main() {
	// N
	var n int
	fmt.Scan(&n)

	// A1 ... An
	a := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a[i], _ = strconv.Atoi(scanner.Text())
	}

	var ans int
	for i := 0; i < n; i++ {
		// 一つ前の集合にa[i]を足せば必ずa[i]が最大になる集合が作れる
		// i=0: 空集合にa[0]を足す
		// i=1: {{}, {a[0]}}にa[1]を足す
		//
		// なのでiの時の集合の数は2^i
		ans = (ans + (a[i]*powpow(2, i))%mod) % mod
	}
	fmt.Println(ans)
}

// powpow caluculate a^b % mod
func powpow(a int, b int) int {
	ret := 1
	i := a
	for b > 0 {
		if b&1 == 1 {
			ret = (ret * i) % mod
		}
		b = b >> 1
		i = (i * i) % mod
	}
	return ret
}

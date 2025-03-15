package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// N S
	var n, s int
	fmt.Scan(&n, &s)
	scanner := bufio.NewScanner(os.Stdin)
	// a1, a2, ..., an
	scanner.Scan()
	line := scanner.Text()
	strs := strings.Split(line, " ")
	a := make([]int, n)
	for i, str := range strs {
		a[i], _ = strconv.Atoi(str)
	}

	solve1000(s, a)
}

func solve500(s int, a []int) {
	ans := false
	for bits := 0; bits < (1 << len(a)); bits += 1 {
		sum := 0
		for i := 0; i < len(a); i++ {
			if (bits>>i)&1 == 1 {
				sum += a[i]
			}
		}
		if sum == s {
			ans = true
			break
		}
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func solve1000(s int, a []int) {
	// dp[i][sum]: カードiまでを用いて、和がsumとなる組み合わせはあるか？
	dp := make([][]bool, len(a)+1)
	for i := 0; i <= len(a); i++ {
		dp[i] = make([]bool, s+1)
	}
	a = append([]int{0}, a...)
	for i := 0; i < len(a); i++ {
		dp[i][0] = true
	}

	for i := 1; i < len(a); i++ {
		for sum := 1; sum <= s; sum++ {
			if dp[i-1][sum] ||
				sum-a[i] >= 0 && dp[i-1][sum-a[i]] {
				dp[i][sum] = true
			}
		}
	}

	if dp[len(a)-1][s] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

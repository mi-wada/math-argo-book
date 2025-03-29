package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	cp := make([]struct{ c, p int }, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		cp[i].c, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		cp[i].p, _ = strconv.Atoi(scanner.Text())
	}

	scanner.Scan()
	q, _ := strconv.Atoi(scanner.Text())
	lr := make([]struct{ l, r int }, q)
	for i := 0; i < q; i++ {
		scanner.Scan()
		lr[i].l, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		lr[i].r, _ = strconv.Atoi(scanner.Text())
	}

	sum := make([][]int, 2)
	sum[0] = make([]int, n+1)
	sum[1] = make([]int, n+1)
	for i := 0; i < n; i++ {
		if cp[i].c == 1 {
			sum[0][i+1] = sum[0][i] + cp[i].p
			sum[1][i+1] = sum[1][i]
		} else {
			sum[0][i+1] = sum[0][i]
			sum[1][i+1] = sum[1][i] + cp[i].p
		}
	}

	for i := 0; i < q; i++ {
		fmt.Println(sum[0][lr[i].r]-sum[0][lr[i].l-1], sum[1][lr[i].r]-sum[1][lr[i].l-1])
	}
}

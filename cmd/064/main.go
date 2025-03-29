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
	// A1 A2 ... AN
	a := make([]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		a[i] = v
	}

	var sum int
	for i := 0; i < n; i++ {
		sum += a[i]
	}

	cond1 := sum <= k
	cond2 := sum%2 == k%2

	if cond1 && cond2 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

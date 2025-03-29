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
	b := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		b[i], _ = strconv.Atoi(scanner.Text())
	}

	var ans int
	ans += b[0]
	ans += b[n-2]
	for i := 0; i < n-2; i++ {
		if b[i] < b[i+1] {
			ans += b[i]
		} else {
			ans += b[i+1]
		}
	}
	fmt.Println(ans)
}

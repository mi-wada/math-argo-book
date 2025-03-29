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
	t := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		scanner.Scan()
		t[i], _ = strconv.Atoi(scanner.Text())
	}
}

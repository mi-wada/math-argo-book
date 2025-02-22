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

	ans := false
	for bits := 0; bits < (1 << n); bits += 1 {
		sum := 0
		for i := 0; i < n; i++ {
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

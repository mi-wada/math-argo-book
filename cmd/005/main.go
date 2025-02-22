package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// N
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	// a1, a2, ..., an
	scanner.Scan()
	line := scanner.Text()
	strs := strings.Split(line, " ")
	a := make([]int, n)
	for i, str := range strs {
		a[i], _ = strconv.Atoi(str)
	}

	sum := 0
	for _, v := range a {
		sum += v
	}
	ans := sum % 100
	fmt.Println(ans)
}

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
	// A1, A2, ..., An
	scanner.Scan()
	line := scanner.Text()
	strs := strings.Split(line, " ")
	a := make([]int, n)
	for i, str := range strs {
		a[i], _ = strconv.Atoi(str)
	}

	ans := 0
	for _, v := range a {
		ans += v
	}
	fmt.Println(ans)
}

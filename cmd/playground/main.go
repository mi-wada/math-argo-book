package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	m, _ := strconv.Atoi(scanner.Text())

	stack := make([]int, 100)
	for i := 0; i < q; i++ {
		scanner.Scan()
		t, _ := strconv.Atoi(scanner.Text())
		switch t {
		case 1:
			scanner.Scan()
			x, _ := strconv.Atoi(scanner.Text())
			stack = append(stack, x)
		case 2:
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Println(last)
		}
	}
}

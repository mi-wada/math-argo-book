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
	a := make([]int, n+1)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 1; i <= n; i++ {
		scanner.Scan()
		s := scanner.Text()
		v, _ := strconv.Atoi(s)
		a[i] = v
	}

	visited := map[int]int{}
	loop := make([]int, 0)
	next := 1
	dist := 0
	for {
		if k == dist {
			fmt.Println(next)
			return
		}
		if visited[next] == 1 {
			loop = append(loop, next)
		}
		if visited[next] == 2 {
			saisho := dist - 2*len(loop)
			remain := k - saisho
			fmt.Println(loop[remain%len(loop)])
			return
		}

		dist++
		visited[next]++
		next = a[next]
	}
}

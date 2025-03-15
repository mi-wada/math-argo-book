package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// n m
	var n, m int
	fmt.Scan(&n, &m)

	// a1 b1
	// ...
	// an bn
	g := make([][]int, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		a--
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		b--
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	visited := make([]bool, n)
	dfs(0, g, visited)

	ans := true
	for _, v := range visited {
		if !v {
			ans = false
			break
		}
	}
	if ans {
		fmt.Println("The graph is connected.")
	} else {
		fmt.Println("The graph is not connected.")
	}
}

func dfs(cur int, g [][]int, visited []bool) {
	if visited[cur] {
		return
	}
	visited[cur] = true
	for _, v := range g[cur] {
		dfs(v, g, visited)
	}
}

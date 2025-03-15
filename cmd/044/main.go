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

	// shortest[i]: 頂点i+1へ行くための最短距離（最小の辺の本数）
	shortest := make([]int, n)
	for i := 0; i < n; i++ {
		shortest[i] = -1
	}

	q := make(queue, 0)
	q.Push(pos{cur: 0, prev: 0})
	for !q.IsEmpty() {
		popped := q.Pop()
		cur, prev := popped.cur, popped.prev
		if shortest[cur] != -1 {
			continue
		}
		shortest[cur] = shortest[prev] + 1
		for _, v := range g[cur] {
			q.Push(pos{
				cur:  v,
				prev: cur,
			})
		}
	}

	for i := 0; i < n; i++ {
		fmt.Println(shortest[i])
	}
}

type pos struct {
	cur  int
	prev int
}

type queue []pos

func (q *queue) Push(x pos) {
	*q = append(*q, x)
}

func (q *queue) Pop() pos {
	popped := (*q)[0]
	*q = (*q)[1:len(*q)]
	return popped
}

func (q queue) IsEmpty() bool {
	return len(q) == 0
}

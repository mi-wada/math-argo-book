package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// N M
	var n, m int
	fmt.Scan(&n, &m)

	ab := make([]struct{ a, b int }, m)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		scanner.Scan()
		ab[i].a, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		ab[i].b, _ = strconv.Atoi(scanner.Text())
	}

	g := make([][]int, n)
	for i := 0; i < m; i++ {
		g[ab[i].a-1] = append(g[ab[i].a-1], ab[i].b-1)
		g[ab[i].b-1] = append(g[ab[i].b-1], ab[i].a-1)
	}

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = -1
	}
	dist[0] = 0
	q := queue{}
	q.Push(0)
	for !q.IsEmpty() {
		cur := q.Pop()
		for _, next := range g[cur] {
			if dist[next] != -1 {
				continue
			}
			dist[next] = dist[cur] + 1
			q.Push(next)
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for _, v := range dist {
		if v > 120 || v == -1 {
			fmt.Fprintln(writer, 120)
		} else {
			fmt.Fprintln(writer, v)
		}
	}
}

type queue []int

func (q *queue) Push(x int) {
	*q = append(*q, x)
}

func (q *queue) Pop() int {
	popped := (*q)[0]
	*q = (*q)[1:len(*q)]
	return popped
}

func (q queue) IsEmpty() bool {
	return len(q) == 0
}

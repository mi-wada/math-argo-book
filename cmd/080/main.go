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

	abc := make([]struct{ a, b, c int }, m)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < m; i++ {
		scanner.Scan()
		abc[i].a, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		abc[i].b, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		abc[i].c, _ = strconv.Atoi(scanner.Text())
	}

	type edge struct {
		next, weight int
	}
	g := make([][]edge, n)
	for i := 0; i < m; i++ {
		g[abc[i].a-1] = append(g[abc[i].a-1], edge{next: abc[i].b - 1, weight: abc[i].c})
		g[abc[i].b-1] = append(g[abc[i].b-1], edge{next: abc[i].a - 1, weight: abc[i].c})
	}

	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = -1
	}
	x[0] = 0
	q := queue{}
	q.Push(0)
	for !q.IsEmpty() {
		cur := q.Pop()
		for _, edge := range g[cur] {
			switch {
			case x[edge.next] == -1:
				q.Push(edge.next)
				x[edge.next] = x[cur] + edge.weight
			// この場合x[edge.next]のままではx[edge.next]は条件を満たさない。なので小さくする必要がある。
			case x[edge.next] > x[cur]+edge.weight:
				q.Push(edge.next)
				x[edge.next] = x[cur] + edge.weight
			// この場合x[edge.next]をこれ以上大きくすることはできないので遷移しない。
			case x[edge.next] == x[cur]+edge.weight:
			// この場合x[edge.next]をこれ以上大きくできない。なぜならば更新したら別のノードから見た制約が崩れる。
			case x[edge.next] < x[cur]+edge.weight:
			}
		}
	}

	fmt.Println(x[n-1])
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

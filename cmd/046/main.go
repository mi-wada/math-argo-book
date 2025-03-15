package main

import (
	"fmt"
)

type cord struct {
	x, y int
}

type pos struct {
	cur, prev cord
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

const (
	dot  = -1
	wall = -2
)

func main() {
	// r c
	var r, c int
	fmt.Scan(&r, &c)
	var s, g cord
	fmt.Scan(&s.y, &s.x)
	s.x--
	s.y--
	fmt.Scan(&g.y, &g.x)
	g.x--
	g.y--

	cells := make([][]int, r)
	for y := 0; y < r; y++ {
		cells[y] = make([]int, c)
	}

	for y := 0; y < r; y++ {
		var s string
		fmt.Scan(&s)
		for x := 0; x < c; x++ {
			switch s[x] {
			case '.':
				cells[y][x] = dot
			case '#':
				cells[y][x] = wall
			default:
				panic(fmt.Sprintf("invalid char: %c", cells[y][x]))
			}
		}
	}

	cells[s.y][s.x] = 0

	q := make(queue, 0)
	q.Push(pos{cur: s, prev: s})
	for !q.IsEmpty() {
		v := q.Pop()
		cur, prev := v.cur, v.prev
		if cells[cur.y][cur.x] != dot && cells[cur.y][cur.x] != 0 {
			continue
		}
		if cells[cur.y][cur.x] != 0 {
			cells[cur.y][cur.x] = cells[prev.y][prev.x] + 1
		}
		for _, next := range []cord{
			{x: cur.x - 1, y: cur.y},
			{x: cur.x + 1, y: cur.y},
			{x: cur.x, y: cur.y - 1},
			{x: cur.x, y: cur.y + 1},
		} {
			if next.x == prev.x && next.y == prev.y {
				continue
			}
			q.Push(
				pos{
					cur:  next,
					prev: cur,
				},
			)
		}
	}
	// for y := 0; y < r; y++ {
	// 	for x := 0; x < c; x++ {
	// 		switch cells[y][x] {
	// 		case dot:
	// 			fmt.Print("  .")
	// 		case wall:
	// 			fmt.Print("  #")
	// 		case 0:
	// 			fmt.Print("  s")
	// 		default:
	// 			fmt.Printf("%3d", cells[y][x])
	// 		}
	// 	}
	// 	fmt.Println()
	// }

	fmt.Println(cells[g.y][g.x])
}

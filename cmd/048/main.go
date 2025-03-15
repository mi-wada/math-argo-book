package main

import (
	"fmt"
	"math"
)

func main() {
	// k
	var k int
	fmt.Scan(&k)

	// dist[i]: mod K == iとなる数値の最小の桁和。dist[0]が答え。
	dist := make([]int, k)
	for i := 0; i < k; i++ {
		dist[i] = math.MaxInt
	}
	dist[1] = 1

	dq := newDeque(k)
	dq.PushBack(1)
	for !dq.Empty() {
		i := dq.PopFront()

		// *10するケース
		// こっちの方がコストが低いので優先的に探索するのでFrontにpush
		next := (i * 10) % k
		if dist[i] < dist[next] {
			dist[next] = dist[i]
			dq.PushFront(next)
		}

		// +1するケース
		next = (i + 1) % k
		if dist[i]+1 < dist[next] {
			dist[next] = dist[i] + 1
			dq.PushBack(next)
		}
	}
	fmt.Println(dist[0])
}

type deque struct {
	s    []int
	from int
	len  int
}

func newDeque(n int) deque {
	s := make([]int, n)
	return deque{s: s, from: 0, len: 0}
}

func (d *deque) PushBack(v int) {
	(*d).s[(d.from+d.len)%len(d.s)] = v
	d.len++
}

func (d *deque) PopBack() int {
	d.len--
	popped := d.s[(d.from+d.len)%len(d.s)]
	return popped
}

func (d *deque) PushFront(v int) {
	d.from = (d.from - 1 + len(d.s)) % len(d.s)
	(*d).s[d.from] = v
	d.len++
}

func (d *deque) PopFront() int {
	d.len--
	popped := d.s[d.from]
	d.from = (d.from + 1) % len(d.s)
	return popped
}

func (d *deque) Empty() bool {
	return d.len == 0
}

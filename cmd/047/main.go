package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	no = iota
	black
	white
)

func main() {
	// n m
	var n, m int
	fmt.Scan(&n, &m)

	// a1 b1
	// ...
	// am bm
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

	colors := make([]int, n)
	for i := 0; i < n; i++ {
		if colors[i] != no {
			continue
		}

		colors[i] = black
		err := dfs(i, colors, g)
		if err != nil {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

func dfs(i int, colors []int, g [][]int) error {
	for _, next := range g[i] {
		switch colors[i] {
		case black:
			switch colors[next] {
			case no:
				colors[next] = white
			case black:
				return errors.New("cannot")
			case white:
				continue
			}
		case white:
			switch colors[next] {
			case no:
				colors[next] = black
			case black:
				continue
			case white:
				return errors.New("cannot")
			}
		default:
			panic("wow")
		}
		err := dfs(next, colors, g)
		if err != nil {
			return err
		}
	}
	return nil
}

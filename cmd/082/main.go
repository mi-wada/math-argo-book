package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	// N
	var n int
	fmt.Scan(&n)

	lr := make(lrs, n)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < n; i++ {
		scanner.Scan()
		lr[i].l, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		lr[i].r, _ = strconv.Atoi(scanner.Text())
	}

	sort.Sort(lr)
	ans := 1
	prevR := lr[0].r
	for _, v := range lr {
		if v.l >= prevR {
			ans++
			prevR = v.r
		}
	}
	fmt.Println(ans)
}

type lr struct {
	l, r int
}

type lrs []lr

var _ sort.Interface = lrs{}

func (s lrs) Len() int {
	return len(s)
}

func (s lrs) Less(i, j int) bool {
	return s[i].r < s[j].r
}

func (s lrs) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

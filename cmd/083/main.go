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

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	a := make(ints, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		a = append(a, v)
	}
	b := make(ints, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		b = append(b, v)
	}

	sort.Sort(a)
	sort.Sort(b)

	var ans int
	for i := 0; i < n; i++ {
		ans += abs(a[i] - b[i])
	}
	fmt.Println(ans)
}

func abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

type ints []int

var _ sort.Interface = ints{}

func (s ints) Len() int {
	return len(s)
}

func (s ints) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s ints) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

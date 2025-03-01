package main

import "fmt"

func main() {
	printMod(2, 56)
}

func printMod(a, b int) {
	fmt.Printf("%d %% %d = %d\n", a, b, a%b)
}

// sum returns the sum of all integers i, from <= i <= to.
func sum(from int, to int) int {
	sum := 0
	for i := from; i <= to; i++ {
		sum += i
	}
	return sum
}

// sumFast returns the sum of all integers i, from <= i <= to.
func sumFast(from int, to int) int {
	return (from + to) * (to - from + 1) / 2
}

func fraction(n int) int {
	if n == 0 {
		return 1
	}
	return n * fraction(n-1)
}

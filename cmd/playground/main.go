package main

import "fmt"

func main() {
	fmt.Printf("sum(1, 100): %v\n", sum(1, 100))
	fmt.Printf("sumFast(1, 100): %v\n", sumFast(1, 100))
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

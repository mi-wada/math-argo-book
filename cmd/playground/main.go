package main

import "fmt"

func main() {
	printMod(1, 1)
	printMod(0, 1)
	printMod(3, 2)
	printMod(-3, 2)
	printMod(-3, -2)
	printMod(3, -2)
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// H W
	var h, w int
	fmt.Scan(&h, &w)

	a := make([][]int, h)
	for i := 0; i < h; i++ {
		a[i] = make([]int, w)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			scanner.Scan()
			v, _ := strconv.Atoi(scanner.Text())
			a[i][j] = v
		}
	}

	rowSums := make([]int, h)
	for i := 0; i < h; i++ {
		var sum int
		for j := 0; j < w; j++ {
			sum += a[i][j]
		}
		rowSums[i] = sum
	}

	colSums := make([]int, w)
	for j := 0; j < w; j++ {
		var sum int
		for i := 0; i < h; i++ {
			sum += a[i][j]
		}
		colSums[j] = sum
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Fprint(writer, rowSums[i]+colSums[j]-a[i][j])
			if j != w-1 {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
}

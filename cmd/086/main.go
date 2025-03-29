package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	cnt := 0
	ok := true
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			cnt++
		} else {
			cnt--
		}
		if cnt < 0 {
			ok = false
		}
	}
	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

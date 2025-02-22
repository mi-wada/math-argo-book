package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		panic(fmt.Sprintf("Example: %s 1000", os.Args[0]))
	}
	binS := os.Args[1]
	decimal, err := binToDec(binS)
	if err != nil {
		panic(err)
	}
	fmt.Printf("binary: %s\n", binS)
	fmt.Printf("decimal: %d\n", decimal)
}

func binToDec(bin string) (int, error) {
	dec := 0
	pow2 := 1
	for i := len(bin) - 1; i >= 0; i-- {
		switch bin[i] {
		case '1':
			dec += pow2
		case '0':
			// do nothing
		default:
			return 0, fmt.Errorf("invalid c: %c", bin[i])
		}
		pow2 *= 2
	}
	return dec, nil
}

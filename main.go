package main

import (
	"fmt"
)

func ackermann(m, n uint64) uint64 {
	if m == 0 {
		return n + 1
	} else if n == 0 {
		return ackermann(m-1, 1)
	} else {
		return ackermann(m-1, ackermann(m, n-1))
	}
}

func main() {
	fmt.Println("__ Ackermann Function's Calculation __")

	for m := 0; m < 5; m++ {
		for n := 0; n < 16-m; n++ {

		}
	}
}

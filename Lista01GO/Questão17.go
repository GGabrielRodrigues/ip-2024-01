package main

import (
	"fmt"
)

var (
	x, y int
)

func main() {
	fmt.Scan(&x, &y)
	if x%2 == 0 {
		fmt.Print(x)
		a := x + 2
		for cont := 1; cont < y; cont++ {
			fmt.Print(" ", a)
			a = a + 2
		}
	} else {
		fmt.Printf("O PRIMEIRO NUMERO NAO E PAR")
	}
}

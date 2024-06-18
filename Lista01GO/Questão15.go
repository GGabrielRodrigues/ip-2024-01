package main

import (
	"fmt"
)

var (
	pares, cont, N int
)

func main() {
	fmt.Scan(&N)
	pares = N / 2
	num := 2
	for cont < pares {
		fmt.Println(num, "^2 = ", num*num)
		cont++
		num = num + 2
	}
}

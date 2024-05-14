package main

import (
	"fmt"
)

var (
	a, b int
)

func main() {
	V := []int{}
	fmt.Scan(&a)
	for m := 1; m <= a; m++ {
		fmt.Scan(&b)
		V = append(V, b)
	}
	fmt.Println(V)
	soma := 0
	for m := 0; m <= a-1; m++ {
		soma = soma + V[m]
	}
	fmt.Println(soma)
}

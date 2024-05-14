package main

import (
	"fmt"
)

var (
	n, b int
)

func main() {
	fmt.Scan(&n)
	a := 1
	V := []int{}
	for a <= n {
		fmt.Scan(&b)
		V = append(V, b)
		a++
	}
	fmt.Println(V)
	inV := []int{}
	for n != 0 {
		b = V[n-1]
		inV = append(inV, b)
		n--
	}
	fmt.Print(inV)
}

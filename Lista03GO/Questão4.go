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
}

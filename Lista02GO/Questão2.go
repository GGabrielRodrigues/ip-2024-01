package main

import (
	"fmt"
)

var (
	A, B float64
	cont int
)

func main() {
	fmt.Scan(&A, &B)
	for cont = 1; cont != 0; cont++ {
		A = A * 1.03
		B = B * 1.015
		if A > B {
			break
		}
	}
	fmt.Printf("ANOS = %v", cont)
}

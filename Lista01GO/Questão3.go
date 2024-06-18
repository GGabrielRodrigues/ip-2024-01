package main

import (
	"fmt"
)

var (
	a, b, c int
)

func main() {
	fmt.Scan(&a, &b, &c)
	if a > 9 || b > 9 || c > 9 {
		fmt.Println("Dígito inválido")
	} else {
		A := a * 100
		B := b * 10
		fmt.Print((A + B + c), ", ", (A+B+c)*(A+B+c))
	}
}

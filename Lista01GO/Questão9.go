package main

import (
	"fmt"
)

var (
	A, B, C float64
)

func main() {
	fmt.Scan(&A, &B, &C)
	fmt.Print("O VALOR DE DELTA E = ", (B*B)-(4*A*C))
}

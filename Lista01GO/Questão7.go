package main

import (
	"fmt"
)

var (
	F, P, mm, C float64
)

func main() {
	fmt.Scan(&F, &P)
	C = (5 * (F - 32)) / 9
	mm = P * 25.4
	fmt.Printf("O VALOR EM CELSIUS = %v\nA QUANTIDADE DE CHUVA E = %v", C, mm)
}

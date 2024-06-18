package main

import (
	"fmt"
)

var (
	nota     float64
	conceito string
)

func main() {
	fmt.Scan(&nota)
	if nota >= 9 && nota <= 10 {
		conceito = "A"
	} else if nota < 9 && nota >= 7.5 {
		conceito = "B"
	} else if nota < 7.5 && nota >= 6 {
		conceito = "C"
	} else if nota < 6 {
		conceito = "D"
	}
	fmt.Printf("NOTA = %v CONCEITO = %v", nota, conceito)
}

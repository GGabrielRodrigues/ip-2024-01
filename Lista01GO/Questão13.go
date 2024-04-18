package main

import (
	"fmt"
)

func main() {

	var nota float64

	fmt.Scan(nota)
	if (nota >= 9) && (nota < 10) {
		fmt.Print("A")
	}
	if (nota >= 7.5) && (nota < 9) {
		fmt.Print("B")
	}

	if (nota >= 6) && (nota < 7.5) {
		fmt.Print("C")
	}
	if (nota >= 0) && (nota < 6) {
		fmt.Print("D")
	}

}

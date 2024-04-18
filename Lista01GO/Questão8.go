package main

import (
	"fmt"
)

func main() {

	var (
		raio, altura, acir, ale, atotal, custo float64
	)

	fmt.Scan(raio)
	fmt.Scan(altura)
	acir = (pi * (raio ^ 2))
	ale = 2 * pi * raio * altura
	atotal = 2*acir + ale
	custo = atotal * 100
	fmt.Print(custo)

}

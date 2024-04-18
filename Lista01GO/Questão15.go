package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		V, ab, h, aresta float64
	)

	fmt.Scan(h, aresta)
	ab = (3 * aresta * aresta * math.Sqrt(3)) / 2
	V = (1 / 3) * ab * h
	fmt.Print("O volume da pirâmide é: ", V)

}

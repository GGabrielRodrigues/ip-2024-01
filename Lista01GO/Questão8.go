package main

import (
	"fmt"
	"math"
)

var (
	R, alt, Ac, Al, At float64
)

func main() {
	pi := math.Pi
	fmt.Scan(&R, &alt)
	Ac = pi * R * R
	Al = 2 * pi * R * alt
	At = 2*Ac + Al
	custo := At * 100
	fmt.Printf("O VALOR DO CUSTO E = %v", custo)
}

package main

import (
	"fmt"
)

var (
	F, C  float64
	tests int
)

func main() {
	sF := []float64{}
	sC := []float64{}
	fmt.Scan(&tests)
	for a := 0; a < tests; a++ {
		fmt.Scan(&F)
		sF = append(sF, F)
		C = (5 * (sF[a] - 32)) / 9
		sC = append(sC, C)
	}
	for a := 0; a < tests; a++ {
		fmt.Printf("%v FAHRENHEIT EQUIVALE(M) A %v CELSIUS\n", sF[a], sC[a])
	}
}

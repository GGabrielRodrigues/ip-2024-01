package main

import (
	"fmt"
)

var (
	somaP, somaI, contP, contI int
)

func main() {
	n := 1
	// p := []int{}
	// i := []int{}
	for n != 0 {
		fmt.Scan(&n)
		if n%2 == 0 && n != 0 {
			// p = append(p,n)
			somaP = n + somaP
			contP++
		} else if n%2 != 0 && n != 0 {
			// i = append(i,n)
			somaI = n + somaI
			contI++
		}
	}
	fmt.Printf("MEDIA PAR: %v\nMEDIA IMPAR: %v", somaP/contP, somaI/contI)
}

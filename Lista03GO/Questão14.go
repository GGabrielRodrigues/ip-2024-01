package main

import (
	"fmt"
)

func main() {

	S := []float64{}
	var qntElem int
	var Indmedfloat, x float64

	fmt.Scan(&qntElem)
	for cont := 0; cont < qntElem; cont++ {
		fmt.Scan(&x)
		S = append(S, x)
	}
	fmt.Printf("-----------------\nLista original:\n%v\n", S)
	fmt.Println("")
	for a := 0; a < qntElem-1; a++ {
		for b := a + 1; b <= qntElem-1; b++ {
			if S[a] > S[b] {
				aux := S[a]
				S[a] = S[b]
				S[b] = aux
			}
		}
	}
	fmt.Printf("Lista ordenada:\n%v\n", S)
	fmt.Println("")
	if len(S)%2 != 0 {
		Indmed := (len(S) / 2)
		fmt.Printf("Mediana = %v", S[Indmed])
	} else {
		Indmed := (len(S) / 2) - 1
		Indmedfloat = (S[Indmed] + S[Indmed+1]) / 2
		fmt.Printf("Mediana = %v", Indmedfloat)
	}

}

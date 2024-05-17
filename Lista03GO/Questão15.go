package main

import (
	"fmt"
)

func main() {
	S := []int{}
	var qntTestes, qntElem, menor, x, numcomp int
	menor = 100000000000

	fmt.Scan(&qntTestes)
	for cont := 0; cont < qntTestes; cont++ {
		fmt.Scan(&qntElem)
		for cont2 := 0; cont2 < qntElem; cont2++ {
			fmt.Scan(&x)
			S = append(S, x)
		}
		for a := 0; a < qntElem; a++ {
			for b := a + 1; b <= qntElem-1; b++ {
				subtr := S[a] - S[b]
				if subtr < 0 {
					subtr = subtr * -1
				}
				fmt.Printf("Subtração: %v\n", subtr)
				if subtr < menor {
					menor = subtr
				}
				numcomp++
			}
		}
		fmt.Printf("Distância entre os elementos mais próximos: %v\nNúmero de comparações: %v", menor, numcomp)
	}

}

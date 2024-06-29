package main

import (
	"fmt"
)

func main() {
	S := []int{}
	var qntElem, x int
	fmt.Scan(&qntElem)
	for cont := 0; cont < qntElem; cont++ {
		fmt.Scan(&x)
		S = append(S, x)
	}
	fmt.Println(S)
	for a := 0; a < qntElem-1; a++ {
		for b := a + 1; b < qntElem; b++ {
			if S[a] > S[b] {
				aux := S[a]
				S[a] = S[b]
				S[b] = aux
			}
		}
	}
	fmt.Printf("%v", S)
}

package main

import (
	"fmt"
)

var (
	a, indMaior, qntNotas, num, maior int
)

func main() {
	fmt.Scan(&qntNotas)
	nota := []int{}
	for cont := 0; cont < qntNotas; cont++ {
		fmt.Scan(&a)
		nota = append(nota, a)
		if nota[cont] > maior {
			maior = nota[cont]
			indMaior = cont
		}
	}
	ult := qntNotas
	fmt.Println(ult)
	freq := make(map[int]int)
	for _, num := range nota {
		freq[num] = freq[num] + 1
	}
	fmt.Printf("Nota %v, %v vezes\n Nota %v, índice %v", nota[qntNotas-1], freq[ult], maior, indMaior)
}

package main

import (
	"fmt"
)

var (
	n int
)

func troca(a int, b int) {
	aux := 0
	aux = a
	a = b
	b = aux
}

func TrocaOpostosSeMenor(vet []int) []int {
	n = len(vet)
	x := 0
	y := n
	for cont := 0; cont < n/2; cont++ {
		if vet[x] < vet[y] {
			troca(vet[x], vet[y])
			// vet[x], vet[y] = vet[y], vet[x]
		}
		x++
		y--
	}
	return vet
}

func main() {
	testes := 0
	tamVet := 0
	a := 0
	vetor := []int{}
	fmt.Scan(&testes)
	for cont := 0; cont < testes; cont++ {
		fmt.Scan(&tamVet)
		for cont2 := 0; cont2 < tamVet; cont2++ {
			fmt.Scan(&a)
			vetor = append(vetor, a)
		}
		fmt.Print(TrocaOpostosSeMenor(vetor))
	}

}

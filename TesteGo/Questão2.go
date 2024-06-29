package main

import (
	"fmt"
)

var (
	n int
)

func troca(a int, b int) {
	a, b = b, a
}

func TrocaOpostosSeMenor(vet []int) []int {
	n = len(vet)
	x := 0
	y := n - 1
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
	vetrandom := []int{-1, -2, 2, 1, -6}
	fmt.Println(vetrandom)
	vetrandom2 := TrocaOpostosSeMenor(vetrandom)
	fmt.Println(vetrandom2)
}

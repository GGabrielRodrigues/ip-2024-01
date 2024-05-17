package main

import (
	"fmt"
)

var (
	n, b int
)

func main() {
	fmt.Scan(&n)
	a, maior := 0, 0
	menor := 100000000000 //100 trilhões
	V := []int{}
	for a < n {
		fmt.Scan(&b)
		V = append(V, b)
		if V[a] > maior {
			maior = V[a]
		}
		if V[a] < menor {
			menor = V[a]
		}
		a++
	}
	fmt.Println(V)
	inV := []int{}
	for n != 0 {
		b = V[n-1]
		inV = append(inV, b)
		n--
	}
	fmt.Print(inV)
	fmt.Printf("\nMaior:%v, menor:%v", maior, menor)
}

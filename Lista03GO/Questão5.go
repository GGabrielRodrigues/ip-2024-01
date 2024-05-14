package main

import (
	"fmt"
	"math/rand"
)

var (
	a, ind int
)

func main() {

	fmt.Scan(&a)
	V := []int{}
	maior := 0
	for m := 1; m <= a; m++ {
		V = append(V, rand.Intn(10))
	}
	fmt.Println(V)
	for m := 0; m <= a-1; m++ {
		if V[m] > maior {
			maior = V[m]
			ind = m + 1
		}
	}
	fmt.Println(ind, maior)

}

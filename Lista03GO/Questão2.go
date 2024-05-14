package main

import (
	"fmt"
)

var (
	N, K, c int
)

func main() {
	fmt.Scan(&N)
	if N >= 1 && N <= 1000 {
		V := []int{}
		for i := 1; i <= N; i++ {
			V = append(V, i)
		}
		fmt.Println(V)
		fmt.Scan(&K)
		for i := 0; i < len(V); i++ {
			if V[i] > K {
				c = c + 1
			}
		}
		fmt.Println(c)
	}

}

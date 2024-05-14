package main

import (
	"fmt"
	"math/rand"
)

var (
	a int
)

func main() {
	a := 1
	for a != 0 {
		fmt.Scan(&a)
		if a != 0 && a > 0 {
			V := []int{}
			maior := 0
			for m := 1; m <= a; m++ {
				V = append(V, rand.Intn(10))
			}
			freq := make(map[int]int)
			for _, num := range V {
				freq[num] = freq[num] + 1
			}
			fmt.Println(V)
			fmt.Println(freq[1])
			for m := 0; m <= a-1; m++ {
				if V[m] > maior {
					maior = V[m]
				}
			}
			for m := 0; m <= maior; m++ {
				fmt.Println("(", m, ")", freq[m])
			}
		}
	}
}

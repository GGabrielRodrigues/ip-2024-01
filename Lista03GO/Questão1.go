package main

import "fmt"

var (
	M int
	B [10]int
)

func main() {
	fmt.Println(10)
	N := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for a := 0; a <= 9; a++ {
		fmt.Printf("%v ", N[a])
	}
	fmt.Println(" ")
	fmt.Scan(&M)
	for a := 0; a < M; a++ {
		fmt.Scan(&B[a])
	}
	for a := 0; a < M; a++ {
		c := 0
		for b := 0; b <= 9; b++ {
			if B[a] == N[b] {
				fmt.Println("ACHEI")
			} else {
				c = c + 1
			}
			if c == 10 {
				fmt.Println("NÃO ACHEI")
			}

		}

	}
}

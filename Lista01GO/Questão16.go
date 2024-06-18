package main

import (
	"fmt"
)

var (
	sal, reaj float64
)

func main() {
	fmt.Scan(&sal)
	if sal <= 300 {
		reaj = sal * 1.5
	} else {
		reaj = sal * 1.3
	}
	fmt.Printf("SALARIO COM REAJUSTE = %v", reaj)
}

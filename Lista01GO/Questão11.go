package main

import (
	"fmt"
)

var (
	num int
)

func main() {
	fmt.Scan(&num)
	if num%3 == 0 && num%5 == 0 {
		fmt.Print("É divisível por 3 e 5")
	} else {
		fmt.Print("Não é divisível por 3 e 5")
	}
}

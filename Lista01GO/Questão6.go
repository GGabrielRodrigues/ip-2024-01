package main

import (
	"fmt"
)

func main() {

	var (
		qntlinhas      int
		faren, celsius float64
	)

	fmt.Scan(qntlinhas)
	for qntlinhas > 0 {
		fmt.Scan(faren)
		celsius = (faren - 32) / 1.8
		fmt.Println(faren, " fahrenheits são iguais a ", celsius, " celsius.")
		qntlinhas = qntlinhas - 1
	}

}

package main

import (
	"fmt"
)

var (
	stacks, H int
)

func main() {
	fmt.Scan(&H)
	if H >= 3 {
		stacks = H / 3
		fmt.Print("O VALOR A PAGAR E = ", stacks*10+10)
	} else {
		fmt.Print("O VALOR A PAGAR E = ", H*5)
	}
}

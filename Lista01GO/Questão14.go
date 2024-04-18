package main

import (
	"fmt"
)

func main() {

	var N, par int

	fmt.Scan(N)
	par = 0
	if (N > 5) && (N < 2000) {
		for par != N {
			par = par + 2
			fmt.Println(par)
		}
	}

}

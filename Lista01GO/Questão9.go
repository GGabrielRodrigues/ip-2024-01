package main

import (
	"fmt"
)

func main() {

	var (
		delta   int
		a, b, c int
	)

	fmt.Scan(a, b, c)
	delta = (b ^ 2 - 4*a*c)
	fmt.Print(delta)

}

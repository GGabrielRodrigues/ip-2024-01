package main

import (
	"fmt"
)

var (
	A1, r, n int
)

func main() {
	fmt.Scan(&A1, &r, &n)
	An := A1 + (n-1)*r
	Sn := ((A1 + An) * n) / 2
	fmt.Print(Sn)
}

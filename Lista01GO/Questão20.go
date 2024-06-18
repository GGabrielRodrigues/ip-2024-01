package main

import (
	"fmt"
)

var (
	h, m, s int
)

func main() {
	fmt.Scan(&h, &m, &s)
	fmt.Print("O TEMPO EM SEGUNDOS E = ", h*3600+m*60+s)
}

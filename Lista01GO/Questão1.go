package main

import (
	"fmt"
)

var (
	nota1, nota2, nota3 float64
)

func main() {
	fmt.Scan(&nota1, &nota2, &nota3)
	fmt.Print("MEDIA = ", ((nota1 + nota2 + nota3) / 3))
}

package main

import (
	"fmt"
)

func main() {

	var (
		far, cel, pol, mm float64
	)

	fmt.Scan(far)
	cel = (far - 32) / 1.8
	fmt.Scan(pol)
	mm = pol * 25.4
	fmt.Println("Valor em celsius: ", cel)
	fmt.Println("Valor em mm: ", mm)

}

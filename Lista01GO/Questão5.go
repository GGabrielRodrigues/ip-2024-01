package main

import (
	"fmt"
)

var (
	numconta         int
	agua, valorconta float64
	tipocon          string
)

func main() {
	fmt.Scan(&numconta, &agua, &tipocon)
	if tipocon == "R" {
		valorconta = 5 + (0.05 * agua)
	} else if tipocon == "C" {
		if agua <= 80 {
			valorconta = 500
		} else {
			valorconta = 500 + (0.25 * (agua - 80))
		}
	} else if tipocon == "I" {
		if agua <= 100 {
			valorconta = 800
		} else {
			valorconta = 800 + (0.04 * (agua - 100))
		}
	}
	fmt.Printf("CONTA = %v\nVALOR DA CONTA = %v", numconta, valorconta)
}

package main

import (
	"fmt"
)

var (
	L, sL, sP, P, NT float64
	mat, freq        int
)

func main() {
	fmt.Scan(&mat)
	for cont := 0; cont < 8; cont++ {
		fmt.Scan(&P)
		sP = sP + P
	}
	MP := sP / 8
	for cont := 0; cont < 5; cont++ {
		fmt.Scan(&L)
		sL = sL + L
	}
	ML := sL / 5
	fmt.Scan(&NT)
	NF := 0.7*MP + 0.15*ML + 0.15*NT
	fmt.Scan(&freq)
	if NF >= 6 && freq >= 96 {
		fmt.Printf("Matrícula: %v, Nota Final: %v, Situação Final: APROVADO", mat, NF)
	} else if NF >= 6 && freq < 96 {
		fmt.Printf("Matrícula: %v, Nota Final: %v, Situação Final: REPROVADO POR FREQUÊNCIA", mat, NF)
	} else if NF < 6 && freq < 96 {
		fmt.Printf("Matrícula: %v, Nota Final: %v, Situação Final: REPROVADO POR NOTA E FREQUÊNCIA", mat, NF)
	} else if NF < 6 && freq >= 96 {
		fmt.Printf("Matrícula: %v, Nota Final: %v, Situação Final: REPROVADO POR NOTA", mat, NF)
	}

}

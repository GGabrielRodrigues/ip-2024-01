package main

import (
	"fmt"
)

var (
	qntElem, a, cont   int // quantidade de elementos do vetor; variável usada para armazenar os valores para depois reinserir no slice//contador genérico
	IndMais, maiorFreq int //Índice do elemento de maior frequência; Frequência do mais frequente; variável que armazena o maior valor
)

func main() {

	S := []int{}           //slice 'S' que armazenará os valores inseridos
	índices := []int{}     //slice 'indices' que armazenará os índices das frequências
	frequências := []int{} //slice 'frequências' que armazenará as frequências de cada elemento do vetor
	maiorFreq := 0
	fmt.Scan(&qntElem)
	for cont = 0; cont < qntElem; cont++ { //loop que acrescenta os valores inseridos pelo usuário ao slice S
		fmt.Scan(&a)
		S = append(S, a)
	}
	fmt.Printf("\n%v", S)

	freq := make(map[int]int) //declaração de um map que fará a contagem da frequência
	for ind, num := range S {
		freq[num] = freq[num] + 1
		índices = append(índices, ind)
		frequências = append(frequências, freq[num])
	}
	for cont = 0; cont < qntElem; cont++ {
		if frequências[cont] >= maiorFreq {
			maiorFreq = frequências[cont]
			IndMais = cont
		}
	}
	fmt.Printf("\n\n%v é o número mais frequente, com uma frequência igual a %v", S[IndMais], maiorFreq)
}

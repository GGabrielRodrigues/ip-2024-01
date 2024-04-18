package main

import (
	"fmt"
)

func main () {

	var (
	pctcadeiras, pctpopular, pctgeral, pctarquibancada, rctcadeiras, rctpopular, rctgeral, rctarquibancada,conta float64
	numtestes,totalpessoas,ordem int
	)
	

	fmt.Print("Quantidade de jogos: ")
	fmt.Scan (numtestes)
	for numtestes > 0
		ordem = 1
		fmt.Println (" ")
		fmt.Println (" ")
		fmt.Print ("Total de pessoas: ")
		fmt.Scan (totalpessoas)
		fmt.Print ("Porcentagem de pessoas na categoria 'Cadeiras': ")
		fmt.Scan (pctcadeiras)
		fmt.Print ("Porcentagem de pessoas na categoria 'Popular': ")
		fmt.Scan (pctpopular)
		fmt.Print ("Porcentagem de pessoas na categoria 'Geral': ")
		fmt.Scan (pctgeral)
		fmt.Print ("Porcentagem de pessoas na categoria 'Arquibancada': ")
		fmt.Scan (pctarquibancada)
		rctcadeiras = (totalpessoas*pctcadeiras/100*20)
		rctgeral = (totalpessoas*pctgeral/100*5)
		rctarquibancada = (totalpessoas*pctarquibancada/100*10)
		rctpopular = (totalpessoas*pctpopular/100)
		conta = rctcadeiras + rctgeral + rctarquibancada + rctpopular
		fmt.Print ("Renda do jogo: ",conta)
		numtestes = numtestes - 1
	}
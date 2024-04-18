package main

import (
	"fmt"
)

var (
	salmin, cadakw, totalkw, valortotal, valordesc float64
)

func main() {

	fmt.Print("Insira o valor do salário mínimo (atual: R$1412): ")
	fmt.Scan(salmin)
	fmt.Print("Insira o valor do total energético gasto, em KW: ")
	fmt.Scan(totalkw)
	cadakw = (salmin * 7 / 1000)
	fmt.Println("Valor de cada KW: ", cadakw)
	valortotal = cadakw * totalkw
	fmt.Println("Valor total a ser pago: R$", valortotal)
	valordesc = valortotal * 9 / 10
	fmt.Print("Valor total com desconto: R$", valordesc)

}

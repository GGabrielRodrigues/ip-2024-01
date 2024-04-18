package main

import (
	"fmt"
)

func main() {

	var horas, valor int

	fmt.Scan(horas)
	if horas >= 3 {
		valor = (int(horas / 3)) * 10
		fmt.Println("Custo: ", valor)
		fmt.Println("Horas: ", horas)
	} else {
		valor = (horas * 5)
		fmt.Println("Custo: ", valor)
		fmt.Println("Horas: ", horas)
	}

}

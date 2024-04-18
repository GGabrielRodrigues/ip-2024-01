package main

import (
	"fmt"
)

func main() {

	var numero int

	fmt.Scan(numero)
	if ((numero % 3) == 0) && ((numero % 5) == 0) {
		fmt.Print("É divisível por 3 e 5")
	} else {
		fmt.Print("NÃO é divisível por 3 e 5")
	}
}

package main

import "fmt"

func main() {
	var nota1, nota2, nota3, media float64
	fmt.Println("Insira a primeira nota: ")
	fmt.Scan(&nota1)
	fmt.Println("Insira a segunda nota: ")
	fmt.Scan(&nota2)
	fmt.Println("Insira a terceira nota: ")
	fmt.Scan(&nota3)
	media = (nota1 + nota2 + nota3) / 3.0
	fmt.Println("A média é: ")
	fmt.Print(media)

}

package main

import "fmt"

var n, input int

func main() {
	fmt.Scanln(&input)

	n = 1

	for n*2 <= input {
		n = n * 2
	}
	fmt.Println(n)
	for n >= 1 {
		fmt.Print(input / n)
		input = input % n
		n = n / 2
	}
}

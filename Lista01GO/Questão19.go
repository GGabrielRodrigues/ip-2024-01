package main

import (
	"fmt"
)

var (
	som, Nu float64
)

func main() {
	fmt.Scan(&Nu)
	cont := 0.1
	som = 1
	for cont = 2; cont != Nu+1; cont++ {
		frac := 1 / cont
		fmt.Printf("1/%v\n", cont)
		som = som + frac
	}
	fmt.Println(cont)
	fmt.Print(som)
}

package main

import (
	"fmt"
)

var (
	numtestes, x, cresc, maiorcresc int
)

func main() {
	fmt.Scan(&numtestes)
	s := []int{}
	for cont := 0; cont < numtestes; cont++ {
		fmt.Scan(&x)
		s = append(s, x)
	}
	fmt.Println(s)
	n := len(s)
	i := 0
	j := i + 1
	for cont := 0; cont < n-1; cont++ {
		if s[i] < s[j] {
			cresc++
			if cresc > maiorcresc {
				maiorcresc = cresc
			}
		} else {
			cresc = 0
		}
		i++
		j++
	}
	fmt.Println("O comprimento do segmento crescente máximo é: ", maiorcresc)
}

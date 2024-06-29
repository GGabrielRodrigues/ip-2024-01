package main

import (
	"fmt"
)

var (
	numtestes, x, iguais, maioriguais int
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
		if s[i] == s[j] {
			iguais++
			if iguais > maioriguais {
				maioriguais = iguais
			}
		} else {
			iguais = 0
		}
		i++
		j++
	}
	fmt.Printf("A maior sequência de elementos iguais possui %v elementos.", maioriguais+1)
}

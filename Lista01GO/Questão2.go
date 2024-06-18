package main

import (
	"fmt"
)

var (
	numtestes            int
	pess, pP, pG, pA, pC float64
)

func main() {

	fmt.Scan(&numtestes)
	for cont := 0; cont < numtestes; cont++ {
		fmt.Scan(&pess, &pP, &pG, &pA, &pC)
		fmt.Println((pess * (pP / 100)) + ((pess * (pG / 100)) * 5) + ((pess * (pA / 100)) * 10) + ((pess * (pC / 100)) * 20))
	}
}

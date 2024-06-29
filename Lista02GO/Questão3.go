package main

import (
	"fmt"
)

var (
	n int
)

func main() {
	fmt.Scan(&n)
	if n == 0 {
		fmt.Print("0! = 1")
	} else {
		for cont := n - 1; cont != 1; cont-- {
			n = n * cont
		}
		fmt.Print(n)
	}
}

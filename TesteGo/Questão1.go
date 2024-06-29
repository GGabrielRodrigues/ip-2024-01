package main

import (
	"fmt"
)

func reverte(str string)(result string) {
	for _,v := range str {
		result = string(v) + result
	}
	return
}

func main() {
 nome := "sahnilab e sotilurip maroda sacnairc"
 nomeINV := reverte(nome)
 fmt.Println(nome)
 fmt.Print(nomeINV)
}
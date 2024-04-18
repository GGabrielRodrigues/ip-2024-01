
 package main

 import (
	"fmt"
 )

func main() {


	var (
	n1,n2,n3,concat int
	quadconcat float64 
	)
	
	fmt.Print ("Digite os três dígitos (redundante, eu sei): ")
	fmt.Scan (n1,n2,n3)
	concat = (n1*100 + n2*10 + n3)
	quadconcat = concat^2
	if (n1 >= 10) || (n2 >= 10) || (n3 >= 10) {
	return fmt.Print ("Dígito Inválido")
	}
	return (
		 fmt.Print( "Concatenação dos dígitos: ",concat)
		 fmt.print ("Quadrado da concatenação: ", quadconcat)
	)


}




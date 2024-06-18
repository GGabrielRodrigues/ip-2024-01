package main

import (
	"fmt"
	"math"
)

var (
	alt, ares, Ab, v float64
)

func main() {
	fmt.Scan(&alt, &ares)
	Ab = (3 * ares * ares * math.Sqrt(3)) / 2
	v = (Ab * alt) / 3
	fmt.Printf("O VOLUME DA PIRAMIDE E = %v METROS CUBICOS", v)
}

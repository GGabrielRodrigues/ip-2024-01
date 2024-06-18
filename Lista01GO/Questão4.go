package main

import (
	"fmt"
)

var (
	sal, totkw float64
)

func main() {
	fmt.Scan(&sal, &totkw)
	valcadakw := ((sal * 0.7) / 100)
	valtodokw := totkw * valcadakw
	valdesc := valtodokw * 0.9
	fmt.Printf("Custo por kW: R$ %v\nCusto do consumo: R$ %v\nCusto com desconto: R$ %v", valcadakw, valtodokw, valdesc)
}

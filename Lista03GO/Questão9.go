package main

import (
	"fmt"
	"math"
)

func dist(a1, b1, c1, a2, b2, c2 float64) float64 {
	x := (a1 - a2) * (a1 - a2)
	y := (b1 - b2) * (b1 - b2)
	z := (c1 - c2) * (c1 - c2)
	F := math.Sqrt(x + y + z)
	return (F)
}

func main() {
	V1 := []float64{}
	V2 := []float64{}
	var i, j, k, r, s, t float64

	fmt.Scan(&i, &j, &k)
	fmt.Scan(&r, &s, &t)
	fmt.Println(dist(i, j, k, r, s, t))

}

package main

import (
	"fmt"
)

var (
	K, conta, s, i, N, incr float32
)

func main() {
	fmt.Scan(&N, &i, &K, &s)
	incr = i + s
	fmt.Printf("Tabuada de soma:\n%v + %v = %v\n", N, i, N+i)
	for conta = 1; conta < K; conta++ {
		fmt.Printf("%v + %v = %v\n", N, incr, N+incr)
		incr = incr + s
	}
	incr = i + s
	fmt.Printf("Tabuada de subtração:\n%v - %v = %v\n", N, i, N-i)
	for conta = 1; conta < K; conta++ {
		fmt.Printf("%v - %v = %v\n", N, incr, N-incr)
		incr = incr + s
	}
	incr = i + s
	fmt.Printf("Tabuada de multiplicação:\n%v x %v = %v\n", N, i, N*i)
	for conta = 1; conta < K; conta++ {
		fmt.Printf("%v x %v = %v\n", N, incr, N*incr)
		incr = incr + s
	}
	incr = i + s
	fmt.Printf("Tabuada de divisão:\n%v / %v = %v\n", N, i, N/i)
	for conta = 1; conta < K; conta++ {
		fmt.Printf("%v / %v = %v\n", N, incr, N/incr)
		incr = incr + s
	}
}

package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndice := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndice] {
				minIndice = j
			}
		}

		arr[i], arr[minIndice] = arr[minIndice], arr[i]
	}
}

func main() {

	desordenado := []int{76, 14, 22, 98, 32, 19, 90}

	fmt.Println("Slice desordenado:", desordenado)

	selectionSort(desordenado)

	fmt.Println("Slice ordenado:", desordenado)
}

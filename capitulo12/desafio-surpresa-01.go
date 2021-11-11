package main

import (
	"fmt"
)

func main() {
	x := somenteImpares(soma, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}...)

	fmt.Println(x)
}

func soma(numeros ...int) int {
	total := 0
	for _, v := range numeros {
		total += v
	}
	return total
}

func somenteImpares(f func(...int) int, numeros ...int) int {
	var impares []int
	for _, v := range numeros {
		if v%2 != 0 {
			impares = append(impares, v)
		}
	}
	total := f(impares...)
	return total
}

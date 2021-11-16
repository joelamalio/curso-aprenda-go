package main

import (
	"fmt"
)

func main() {
	x := []int{5, 10, 15, 20}
	fmt.Println(soma1(x...))
	fmt.Println(soma2(x))
}

func soma1(inteiros ...int) int {
	soma := 0
	for _, v := range inteiros {
		soma += v
	}
	return soma
}

func soma2(inteiros []int) int {
	soma := 0
	for _, v := range inteiros {
		soma += v
	}
	return soma
}

package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for i, value := range slice {
		fmt.Println(i, "-", value)
	}

	fmt.Printf("Tipo da estrutura: %T", slice)
}

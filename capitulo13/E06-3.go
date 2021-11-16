package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Linha 1")
	fmt.Println("Linha 2")
	fmt.Println("Linha 3")
	defer fmt.Println("Linha 4")
	fmt.Println("Linha 5")
}

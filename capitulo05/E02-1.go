/*
- Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/

package main

import (
	"fmt"
)

func main() {
	x := 22
	fmt.Printf("decimal: %d\n", x)
	fmt.Printf("binario: %b\n", x)
	fmt.Printf("hexadecimal: %#x", x)
}

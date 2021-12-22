/*
- Nível 10?! Êita! Parabéns!
- Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
    - Usando uma função anônima auto-executável
    - Usando buffer
*/

package main

import (
	"fmt"
)

func main() {
	usandoFuncaoAnonima()
	fmt.Println("")
	usandoBuffer()
}

func usandoFuncaoAnonima() {
	fmt.Println("Usando uma função anônima auto-executável")
	c := make(chan int)

	go func(c chan int) {
		c <- 42
	}(c)

	fmt.Println(<-c)
}

func usandoBuffer() {
	fmt.Println("Usando buffer")
	c := make(chan int, 1)

	c <- 1

	fmt.Println(<-c)
}

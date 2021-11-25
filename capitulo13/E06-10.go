/*
- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/

package main

import (
	"fmt"
)

func main() {
	c1 := counter()
	c2 := counter()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())

}

func counter() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

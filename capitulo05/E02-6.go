/*
- Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
- Demonstre estes valores.
*/

package main

import (
	"fmt"
)

const (
	_ = 2048 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Printf("%v\n%v\n%v\n%v", a, b, c, d)
}

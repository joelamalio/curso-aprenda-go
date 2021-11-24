/*
- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
*/

package main

import (
	"fmt"
)

func main() {
	a := `teste
	raw		string
		literal`
	fmt.Println(a)
}

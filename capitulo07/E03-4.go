/*
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import (
	"fmt"
)

func main() {

	ano := 1986

	for {
		if ano > 2021 {
			break
		}
		fmt.Println(ano)
		ano++
	}
}

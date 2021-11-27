/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/

package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "futebol"

	switch esporteFavorito {
	case "futebol":
		fmt.Printf("Você gosta de jogar com os pés")
	case "basket":
		fmt.Printf("Você gosta de jogar com as mãos")
	case "surf":
		fmt.Printf("Você gosta de surfar")

	}
}

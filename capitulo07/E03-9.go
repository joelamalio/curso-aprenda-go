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

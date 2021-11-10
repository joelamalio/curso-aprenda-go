package main

import (
	"fmt"
)

func main() {
	pessoas := [][]string{

		{
			"Nome",
			"Sobrenome",
			"Hobby favorito",
		},
		{
			"Tony",
			"Stark",
			"Voar na cidade",
		},
		{
			"Peter",
			"Parker",
			"Saltar de um prédio para o outro",
		},
		{
			"Capitão",
			"América",
			"Defender a ordem",
		},
	}

	for _, pessoa := range pessoas {
		for _, dados := range pessoa {
			fmt.Printf("%v\t", dados)
		}
		fmt.Printf("\n")
	}
}

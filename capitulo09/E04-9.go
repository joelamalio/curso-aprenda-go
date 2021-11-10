package main

import (
	"fmt"
)

func main() {
	mapa := map[string][]string{
		"stark_tony":      {"Voar na cidade", "Disparar raios de energia", "Falar com Jarvis"},
		"parker_peter":    {"Saltar de um prédio para o outro", "Soltar teias", "Escalar paredes"},
		"américa_capitão": {"Defender a ordem", "Lançar o escudo"},
	}

	mapa["incrível_hulk"] = []string{"Saltar", "Gritar", "Destruir"}

	for chave, valor := range mapa {
		fmt.Println(chave)
		for _, hobbie := range valor {
			fmt.Printf("\t%v\n", hobbie)
		}
	}
}

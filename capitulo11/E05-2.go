package main

import (
	"fmt"
)

type pessoa struct {
	nome                      string
	sobrenome                 string
	saboresFavoritosDeSorvete []string
}

func main() {
	pessoa1 := pessoa{
		nome:                      "João",
		sobrenome:                 "da Silva",
		saboresFavoritosDeSorvete: []string{"chocolate", "coco", "manga"},
	}

	pessoa2 := pessoa{
		nome:                      "Maria",
		sobrenome:                 "Pereira",
		saboresFavoritosDeSorvete: []string{"Pitanga", "umbu"},
	}

	pessoas := map[string]pessoa{
		pessoa1.sobrenome: pessoa1,
		pessoa2.sobrenome: pessoa2,
	}

	for _, p := range pessoas {
		fmt.Println(p.nome, pessoa2.sobrenome)
		for _, s := range pessoa2.saboresFavoritosDeSorvete {
			fmt.Printf("\t%v\n", s)
		}
	}
}

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

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, s := range pessoa1.saboresFavoritosDeSorvete {
		fmt.Printf("\t%v\n", s)
	}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, s := range pessoa2.saboresFavoritosDeSorvete {
		fmt.Printf("\t%v\n", s)

	}
}

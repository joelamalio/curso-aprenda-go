package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (p pessoa) dadoscompletos() string {
	return fmt.Sprintf("nome: %v %v\nidade: %d", p.nome, p.sobrenome, p.idade)
}

func main() {
	pessoa := pessoa{
		nome:      "João",
		sobrenome: "das Botas",
		idade:     60,
	}

	fmt.Println(pessoa.dadoscompletos())
}
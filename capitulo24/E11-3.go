/*
- Crie um struct "erroEspecial" que implemente a interface builtin.error.
- Crie uma função que tenha um valor do tipo error como parâmetro.
- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.
*/

package main

import "fmt"

func main() {
	es := erroEspecial{
		codigo:   1000,
		mensagem: "Deu ruim no processamento",
	}

	processaErro(es)
}

type erroEspecial struct {
	codigo   int
	mensagem string
}

func (es erroEspecial) Error() string {
	return fmt.Sprintf("%v - %v", es.codigo, es.mensagem)
}

func processaErro(err error) {
	fmt.Println(err.Error())
}

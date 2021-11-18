package main

import (
	"fmt"
)

func retornaFuncao() func() string {
	return func() string {
		return "Oi"
	}
}

func main() {
	funcao := retornaFuncao()
	fmt.Println(funcao())
}

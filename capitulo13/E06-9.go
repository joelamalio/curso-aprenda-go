package main

import (
	"fmt"
)

func main() {
	executarFuncao(hello)
}

func hello() {
	fmt.Println("Hello, playground")
}

func executarFuncao(funcao func()) {
	funcao()
}

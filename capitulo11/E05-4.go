package main

import (
	"fmt"
)

func main() {
	anonimo := struct {
		numeros map[int]string
		cores   []string
	}{
		numeros: map[int]string{
			1:  "um",
			2:  "dois",
			10: "dez",
		},
		cores: []string{"azul", "vermelho", "branco"},
	}

	fmt.Println(anonimo)
}

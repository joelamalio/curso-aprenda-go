package main

import (
	"fmt"
)

func main() {
	var nota float64 = 6

	switch {
	case nota == 10:
		fmt.Printf("Parabéns!!! Você é muito bom! Tirou a nota máxima e foi aprovado")
	case nota >= 6:
		fmt.Printf("Parabéns!!! Você foi aprovado com a nota %f", nota)
	default:
		fmt.Printf("Você foi reprovado com a nota %f", nota)

	}
}

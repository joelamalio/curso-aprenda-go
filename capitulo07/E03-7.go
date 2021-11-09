package main

import (
	"fmt"
)

func main() {
	var nota float64 = 10

	if nota == 10 {
		fmt.Printf("Parabéns!!! Você é muito bom! Tirou a nota máxima e foi aprovado")
	} else if nota >= 6 {
		fmt.Printf("Parabéns!!! Você foi aprovado com a nota %f", nota)
	} else {
		fmt.Printf("Você foi reprovado com a nota %f", nota)
	}
}

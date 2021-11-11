package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	caminhonete := caminhonete{
		veiculo: veiculo{
			portas: 2,
			cor:    "verde",
		},
		tracaoNasQuatro: true,
	}

	sedan := sedan{
		veiculo: veiculo{
			portas: 4,
			cor:    "azul",
		},
		modeloLuxo: true,
	}

	fmt.Println(caminhonete)
	fmt.Println(sedan)
	fmt.Println(caminhonete.portas)
	fmt.Println(sedan.veiculo.cor)
}

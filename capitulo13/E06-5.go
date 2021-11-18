package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return 2 * math.Pi * c.raio
}

type figura interface {
	area() float64
}

func info(f figura) float64 {
	return f.area()
}

func main() {
	q := quadrado{
		lado: 10,
	}

	c := circulo{
		raio: 100,
	}

	fmt.Println("Área do quadrado:", info(q))
	fmt.Println("Área do círculo:", info(c))
}

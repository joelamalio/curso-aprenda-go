/*
- Nos capítulos seguintes, uma das coisas que veremos é testes.
- Para testar sua habilidade de se virar por conta própria... desafio:
    - Utilizando as seguintes fontes: https://godoc.org/testing & https://www.golang-book.com/books/intro/12
    - Tente descobrir por conta própria como funcionam testes em Go.
    - Pode usar tradutor automático, pode rodar código na sua máquina, pode procurar no Google. Vale tudo.
    - O exercício é: crie um teste simples de uma função ou método ou pedaço qualquer de código.
*/

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	sum := SumAll(numbers...)
	fmt.Println(sum)
}

func SumAll(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

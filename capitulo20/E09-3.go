/*
- Utilizando goroutines, crie um programa incrementador:
    - Tenha uma variável com o valor da contagem
    - Crie um monte de goroutines, onde cada uma deve:
        - Ler o valor do contador
        - Salvar este valor
        - Fazer yield da thread com runtime.Gosched()
        - Incrementar o valor salvo
        - Copiar o novo valor para a variável inicial
    - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
    - Demonstre que há uma condição de corrida utilizando a flag -race
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	numeroGoroutines := 100
	contador := 0

	wg.Add(numeroGoroutines)

	for i := 0; i < numeroGoroutines; i++ {
		go func(x int) {
			valor := x
			runtime.Gosched()
			valor++
			contador = valor
			wg.Done()
		}(contador)
	}

	fmt.Println("Numero de Goroutines:", numeroGoroutines)
	fmt.Println("Contador:", contador)

	wg.Wait()
}

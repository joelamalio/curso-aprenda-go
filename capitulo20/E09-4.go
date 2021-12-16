/*
- Utilize mutex para consertar a condição de corrida do exercício anterior.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex
var contador int

func main() {
	numeroGoroutines := 100

	wg.Add(numeroGoroutines)

	for i := 0; i < numeroGoroutines; i++ {
		go func() {
			mu.Lock()
			valor := contador
			runtime.Gosched()
			valor++
			contador = valor
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Numero de Goroutines:", numeroGoroutines)
	fmt.Println("Contador:", contador)
}

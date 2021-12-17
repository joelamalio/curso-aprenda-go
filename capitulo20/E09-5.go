/*
- Utilize atomic para consertar a condição de corrida do exercício #3.
*/

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	numeroGoroutines := 100
	var contador int32 = 0

	wg.Add(numeroGoroutines)

	for i := 0; i < numeroGoroutines; i++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("Numero de Goroutines:", numeroGoroutines)
	fmt.Println("Contador:", contador)
}

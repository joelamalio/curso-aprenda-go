/*
- Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	go imprimenumero(1)
	go imprimenumero(2)

	wg.Wait()
}

func imprimenumero(numero int) {
	fmt.Println("Iniciando a função", numero)
	wg.Done()
}

/*
- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
- Tire estes números do canal e demonstre-os.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int)

	write(channel)

	read(channel)

	wg.Wait()

	close(channel)

}

func write(channel chan<- int) {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 10; j++ {
				channel <- j
			}
			wg.Done()
		}()
	}
}

func read(channel <-chan int) {
	go func() {
		for {
			value := <-channel
			fmt.Println(value)
		}
	}()
}

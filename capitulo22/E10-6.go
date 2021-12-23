/*
- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
*/

package main

import "fmt"

func main() {
	channel := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			channel <- i
		}
		close(channel)
	}()

	for value := range channel {
		fmt.Println(value)
	}
}

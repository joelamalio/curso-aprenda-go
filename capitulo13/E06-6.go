package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Executando uma função anônima")
	}()
}

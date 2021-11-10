package main

import (
	"fmt"
)

func main() {
	array := [5]int{}
	array[0] = 10
	array[1] = 20
	array[2] = 30
	array[3] = 40
	array[4] = 50

	for i, value := range array {
		fmt.Println(i, "-", value)
	}

	fmt.Printf("Tipo do array: %T", array)
}

package main

import (
	"fmt"
)

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Println("Do primeiro ao terceiro item do slice")
	for i, value := range slice[:3] {
		fmt.Println(i, "-", value)
	}

	fmt.Println("Do quinto ao último item do slice")
	for i, value := range slice[4:] {
		fmt.Println(i, "-", value)
	}

	fmt.Println("Do segundo ao sétimo item do slice")
	for i, value := range slice[1:7] {
		fmt.Println(i, "-", value)
	}

	fmt.Println("Do terceiro ao penúltimo item do slice")
	for i, value := range slice[2:9] {
		fmt.Println(i, "-", value)
	}

	fmt.Println("Do terceiro ao penúltimo item do slice usando a função len()")
	for i, value := range slice[2:(len(slice) - 1)] {
		fmt.Println(i, "-", value)
	}
}

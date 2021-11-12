package main

import (
	"fmt"
)

func main() {
	fmt.Println(retornaint())
	fmt.Println(retornaintestring())
}

func retornaint() int {
	return 10
}

func retornaintestring() (int, string) {
	return 20, "Bom dia!"
}

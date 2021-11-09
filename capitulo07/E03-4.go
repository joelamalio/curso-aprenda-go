package main

import (
	"fmt"
)

func main() {

	ano := 1986

	for {
		if ano > 2021 {
			break
		}
		fmt.Println(ano)
		ano++
	}
}

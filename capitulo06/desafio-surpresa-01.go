package main

import (
	"fmt"
)

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%d\t%#x\t%c\n", i, i, i)
	}
}

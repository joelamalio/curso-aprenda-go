package main

import (
	"fmt"
)

const (
	_ = 2048 + iota
	a
	b
	c
	d
)

func main() {
	fmt.Printf("%v\n%v\n%v\n%v", a, b, c, d)
}

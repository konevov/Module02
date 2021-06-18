package main

import (
	"fmt"
)

func main() {
	var A *int
	B := 88
	A = &B
	fmt.Println(" A = ", *A)
	*A = 32
	fmt.Println(" B = ", B)
}

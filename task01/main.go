package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "104"
	num := 35

	strToNum, _ := strconv.Atoi(str)
	numToStr := strconv.Itoa(num)
	fmt.Printf("Type str - %T, Type strToNum - %T\n", str, strToNum)
	fmt.Printf("Type num - %T, Type numTostr - %T\n", num, numToStr)
}

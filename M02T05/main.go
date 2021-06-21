package main

import (
	"fmt"
)

func contains(a []string, x string) bool {
	for _, v := range a {
		if x == v {
			return true
		}
	}
	return false
}

func getMax(v ...int64) (res int64) {
	res = -1 << 63
	for _, v := range v {
		if v > res {
			res = v
		}
	}
	return res
}

func main() {
	a := []string{"11", "22", "33"}
	x := "33"
	fmt.Println("Result of the contains function:", contains(a, x))
	fmt.Println("Result of the getMax function:", getMax(-2, -4))
}

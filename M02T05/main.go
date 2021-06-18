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

func getMax(v ...int) (res int) {
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
	fmt.Println("Result of the getMax function:", getMax(1, 2, 3, 4, 5, 23, 7, 6))

}

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

func getMax(v ...int64) (res int64, err string) {
	if v == nil {
		err = "Error - no data"
		return
	}

	res = v[0]
	for _, v := range v {

		if v > res {
			res = v
		}
	}
	return
}

func main() {
	a := []string{"11", "22", "33"}
	x := "33"
	fmt.Println("Result of the contains function:", contains(a, x))

	max, err := getMax(-2)
	if err != "" {
		fmt.Println("Result of the getMax function:", err)
		return
	}
	fmt.Println("Result of the getMax function:", max)
}

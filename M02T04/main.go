package main

import "fmt"

func main() {
	bookBase := map[string]map[string][]string{
		"John": {
			"book": {"Go in Action", "Dune", "Star wars"},
		},
		"Bob": {
			"periodicals": {"Times", "Forbes"},
		},
		"Tom": {
			"book": {"Other"},
		},
		"Greg":   {},
		"Bjarne": {},
	}

	counter := 0
	for key, val := range bookBase {

		for _, v := range val {
			counter++
			num := len(v)
			fmt.Printf("Имя: %s, количество изданий на руках: %v\n", key, num)
		}

	}
	fmt.Println("Количество читателей с изданиями на руках:", counter)
}

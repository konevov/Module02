package main

import (
	"fmt"
)

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
	for key1 := range bookBase {
		key2 := bookBase[key1]

		for _, v := range key2 {
			counter++
			num := len(v)
			fmt.Printf("Имя: %s, количество изданий на руках: %v\n", key1, num)
		}

	}
	fmt.Println("Количество читателей с изданиями на руках:", counter)
	fmt.Println("Всего читателей:", len(bookBase))
}

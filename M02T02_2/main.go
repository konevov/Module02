package main

import (
	"fmt"
	"math"
)

func main() {
	R := new(float64)
	*R = 35 / (2 * math.Pi)
	S := math.Pi * *R * *R
	fmt.Printf(" R круга  = %.2f\n S круга = %.2f\n", (math.Round(*R*100))/100, (math.Round(S*100))/100)

}

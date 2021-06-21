package main

import (
	"fmt"
	"math"
)

type AmericanVelocity float32
type EuropeanVelocity float32

func main() {

	var dataA AmericanVelocity = 120.4 * 3.6
	var dataE EuropeanVelocity = 130 * 2.237
	fmt.Printf("American velocity - %.2f\n", math.Round(float64(dataA)*100)/100)
	fmt.Printf("European velocity - %.2f\n", math.Round(float64(dataE)*100)/100)
}

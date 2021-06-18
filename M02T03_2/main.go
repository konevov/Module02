package main

import (
	"fmt"
)

func main() {
	workingDays := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	weekEnd := []string{"Saturday", "Sunday"}
	fullWeek := append(workingDays, weekEnd...)
	fmt.Printf("Full Week%s\n", fullWeek)
}

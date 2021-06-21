package main

import (
	"fmt"
)

func main() {
	week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	workingDays := week[0:5]
	week = week[5:7]
	fmt.Println("Working days -", workingDays)
	fmt.Println("Weekend", week)
}

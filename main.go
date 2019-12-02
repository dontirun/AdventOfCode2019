package main

import (
	"fmt"

	day1 "./day1"
	day2 "./day2"
)

var days = []func(){day1.Solutions, day2.Solutions}

func main() {
	for day, solution := range days {
		fmt.Printf("Details for Day %d: \n", day+1)
		solution()
	}
}

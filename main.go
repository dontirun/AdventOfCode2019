package main

import (
	"fmt"

	day1 "./day1"
)

var days = []func(){day1.Solutions}

func main() {
	for day, solution := range days {
		fmt.Printf("Details for Day %d: \n", day+1)
		solution()
	}
}

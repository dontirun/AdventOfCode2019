package day4

import (
	"fmt"
	"time"
)

var parts = []func(){partOne, partTwo}

//Solutions prints the solution and runtime for Day 4 Solutions
func Solutions() {
	for i, fn := range parts {
		fmt.Printf("\nStarting Part %d \n", i+1)
		start := time.Now()
		fn()
		fmt.Printf("Completed Part %d, took %s\n", i+1, time.Since(start))
	}
}

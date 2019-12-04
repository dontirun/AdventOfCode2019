package day4

import (
	"fmt"
)

//Prints the solution for Day 4 Part 2
func partTwo() {
	lowerBoundsSlice, upperBoundsSlice := parseProgramValues()
	for key := range lowerBoundsSlice {
		fmt.Printf("%d\n", computeValidPasswords2(lowerBoundsSlice[key], upperBoundsSlice[key]))
	}
}

func computeValidPasswords2(lowerBound int, upperBound int) (total int) {
	total = 0
	for num := lowerBound; num <= upperBound; num++ {
		samenessCounter := 0
		onlyTwo := false
		for intLength := 0; intLength < 6; intLength++ {
			samenessCheck, decreaseCheck := checkAdjacentModulo(num, intLength+1, intLength)
			if !decreaseCheck {
				onlyTwo = false
				break
			}
			if samenessCheck {
				samenessCounter++
			} else if samenessCounter == 1 {
				onlyTwo = true
			} else {
				samenessCounter = 0
			}
		}
		if onlyTwo {
			total++
		}
	}
	return
}

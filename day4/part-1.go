package day4

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//Prints the solution for Day 4 Part 1
func partOne() {
	lowerBoundsSlice, upperBoundsSlice := parseProgramValues()
	for key := range lowerBoundsSlice {
		fmt.Printf("%d \n", computeValidPasswords(lowerBoundsSlice[key], upperBoundsSlice[key]))
	}
}

func parseProgramValues() (lowerBoundsSlice []int, upperBoundsSlice []int) {
	file, err := os.Open("./inputs/d4-p1-p2")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "-")
		lowerBounds, _ := strconv.Atoi(splitLine[0])
		upperBounds, _ := strconv.Atoi(splitLine[1])
		lowerBoundsSlice = append(lowerBoundsSlice, lowerBounds)
		upperBoundsSlice = append(upperBoundsSlice, upperBounds)
	}
	return
}

func computeValidPasswords(lowerBound int, upperBound int) (total int) {
	total = 0
	for num := lowerBound; num <= upperBound; num++ {
		samenessCounter := 0
		for intLength := 0; intLength < 6; intLength++ {
			samenessCheck, decreaseCheck := checkAdjacentModulo(num, intLength+1, intLength)
			if !decreaseCheck {
				samenessCounter = -1
				break
			}
			if samenessCheck {
				samenessCounter++
			}
		}
		if samenessCounter >= 1 {
			total++
		}
	}
	return
}

func checkAdjacentModulo(num int, leftDigitPosition int, rightDigitPosition int) (samenessCheck bool, decreaseCheck bool) {
	value := (num/(int(math.Pow10(rightDigitPosition)))%10 - (num/int(math.Pow10(leftDigitPosition)))%10)
	switch {
	case value == 0:
		samenessCheck = true
		decreaseCheck = true
		return
	case value > 0:
		samenessCheck = false
		decreaseCheck = true
		return
	case value < 0:
		samenessCheck = false
		decreaseCheck = false
		return
	default:
		log.Fatalf("Shouldn't reach here")
	}
	return
}

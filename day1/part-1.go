package day1

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

//Prints the solution for Day 1 Part 1
func partOne() {
	massList, err := os.Open("./inputs/d1-p1")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	scanner := bufio.NewScanner(massList)
	totalFuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalf("failed parsing intger from string: %s", err)
		}
		totalFuel += calculateFuelRequirement(mass)
	}
	massList.Close()
	println(totalFuel)
}

// Calculates fuel needed for individual modules to launch.
func calculateFuelRequirement(mass int) (fuelRequired int) {
	return mass/3 - 2
}

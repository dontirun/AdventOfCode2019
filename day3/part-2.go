package day3

import (
	"fmt"
	"log"
	"strconv"
)

type coordInfo struct {
	touched    int
	totalSteps int
}

//Prints the solution for Day 3 Part 2
func partTwo() {
	allWiresMap := map[string]coordInfo{}
	minSteps := -1
	allWireDirections := parseProgramValues()
	totalWires := len(allWireDirections)
	for wireNum := 0; wireNum < totalWires; wireNum++ {
		minSteps = updateGrid2(allWiresMap, allWireDirections[wireNum], wireNum+1 == totalWires, totalWires, minSteps)
	}
	if minSteps != -1 {
		fmt.Printf("%d \n", minSteps)
	} else {
		fmt.Println("All the wires did not intersect")
	}
}

//Updates the allWiresMap and returns the minimum steps required given a specific wire path
func updateGrid2(allWiresMap map[string]coordInfo, wireDirections []string, isLastWire bool, totalWires int, minSteps int) int {
	currentX, currentY := 0, 0
	totalSteps := 0
	currentWireMap := map[string]bool{}
	for _, direction := range wireDirections {
		length, _ := strconv.Atoi(direction[1:])
		switch {
		case string(direction[0]) == "U":
			for dist := 1; dist <= length; dist++ {
				totalSteps++
				mapString := fmt.Sprintf("%d,%d", currentX, currentY+dist)
				minSteps = updateGrid2Helper(allWiresMap, mapString, totalSteps, isLastWire, totalWires, currentWireMap, minSteps)
			}
			currentY += length
			break
		case string(direction[0]) == "D":
			for dist := 1; dist <= length; dist++ {
				totalSteps++
				mapString := fmt.Sprintf("%d,%d", currentX, currentY-dist)
				minSteps = updateGrid2Helper(allWiresMap, mapString, totalSteps, isLastWire, totalWires, currentWireMap, minSteps)
			}
			currentY -= length
			break
		case string(direction[0]) == "L":
			for dist := 1; dist <= length; dist++ {
				totalSteps++
				mapString := fmt.Sprintf("%d,%d", currentX-dist, currentY)
				minSteps = updateGrid2Helper(allWiresMap, mapString, totalSteps, isLastWire, totalWires, currentWireMap, minSteps)
			}
			currentX -= length
			break
		case string(direction[0]) == "R":
			for dist := 1; dist <= length; dist++ {
				totalSteps++
				mapString := fmt.Sprintf("%d,%d", currentX+dist, currentY)
				minSteps = updateGrid2Helper(allWiresMap, mapString, totalSteps, isLastWire, totalWires, currentWireMap, minSteps)
			}
			currentX += length
			break
		default:
			log.Fatalf("Should not reach here")
		}
	}
	return minSteps
}

//Function for redundent logic in updateGridHelper switch statements
func updateGrid2Helper(allWiresMap map[string]coordInfo, mapString string, totalSteps int, isLastWire bool, totalWires int, currentWireMap map[string]bool, minSteps int) int {
	if !currentWireMap[mapString] {
		if val, ok := allWiresMap[mapString]; ok {
			allWiresMap[mapString] = coordInfo{touched: val.touched + 1, totalSteps: val.totalSteps + totalSteps}
			if isLastWire && val.touched+1 == totalWires {
				minSteps = findMinSteps(minSteps, allWiresMap[mapString].totalSteps)
			}
		} else {
			allWiresMap[mapString] = coordInfo{touched: val.touched + 1, totalSteps: totalSteps}
		}
	}
	currentWireMap[mapString] = true
	return minSteps
}

//Returns the smaller step value of the current intersection and the current smallest
func findMinSteps(currentMinSteps int, comparedValue int) int {
	if currentMinSteps == -1 || comparedValue < currentMinSteps {
		return comparedValue
	}
	return currentMinSteps
}

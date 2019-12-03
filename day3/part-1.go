package day3

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//Prints the solution for Day 3 Part 1
func partOne() {
	allWiresMap := map[string]int{}
	manhattanDist := -1
	allWireDirections := parseProgramValues()
	for _, wire := range allWireDirections {
		manhattanDist = updateGrid(allWiresMap, wire, manhattanDist)
	}
	if manhattanDist != -1 {
		fmt.Printf("%d \n", manhattanDist)
	} else {
		fmt.Println("All the wires did not intersect")
	}
}

func parseProgramValues() (allWireDirections [][]string) {
	file, err := os.Open("./inputs/d3-p1-p2")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wireString := scanner.Text()
		wireDirections := strings.Split(wireString, ",")
		allWireDirections = append(allWireDirections, wireDirections)
	}
	return
}

//Updates the allWiresMap and returns the smallest manhattan distance for wire intersections
func updateGrid(allWiresMap map[string]int, wireDirections []string, manhattanDist int) int {
	currentX, currentY := 0, 0
	currentWireMap := map[string]bool{}
	for _, direction := range wireDirections {
		switch {
		case string(direction[0]) == "U":
			length, _ := strconv.Atoi(direction[1:])
			for dist := 1; dist <= length; dist++ {
				mapString := fmt.Sprintf("%d,%d", currentX, currentY+dist)
				if val, ok := allWiresMap[mapString]; ok {
					allWiresMap[mapString] = val + 1
					if !currentWireMap[mapString] {
						manhattanDist = minManhattan(manhattanDist, currentX, currentY+dist)
					}
				} else {
					allWiresMap[mapString] = 1
				}
				currentWireMap[mapString] = true
			}
			currentY += length
			break
		case string(direction[0]) == "D":
			length, _ := strconv.Atoi(direction[1:])
			for dist := 1; dist <= length; dist++ {
				mapString := fmt.Sprintf("%d,%d", currentX, currentY-dist)
				if val, ok := allWiresMap[mapString]; ok {
					allWiresMap[mapString] = val + 1
					if !currentWireMap[mapString] {
						manhattanDist = minManhattan(manhattanDist, currentX, currentY-dist)
					}

				} else {
					allWiresMap[mapString] = 1
				}
				currentWireMap[mapString] = true
			}
			currentY -= length
			break
		case string(direction[0]) == "L":
			length, _ := strconv.Atoi(direction[1:])
			for dist := 1; dist <= length; dist++ {
				mapString := fmt.Sprintf("%d,%d", currentX-dist, currentY)
				if val, ok := allWiresMap[mapString]; ok {
					allWiresMap[mapString] = val + 1
					if !currentWireMap[mapString] {
						manhattanDist = minManhattan(manhattanDist, currentX-dist, currentY)
					}
				} else {
					allWiresMap[mapString] = 1
				}
				currentWireMap[mapString] = true
			}
			currentX -= length
			break
		case string(direction[0]) == "R":
			length, _ := strconv.Atoi(direction[1:])
			for dist := 1; dist <= length; dist++ {
				mapString := fmt.Sprintf("%d,%d", currentX+dist, currentY)
				if val, ok := allWiresMap[mapString]; ok {
					allWiresMap[mapString] = val + 1
					if !currentWireMap[mapString] {
						manhattanDist = minManhattan(manhattanDist, currentX+dist, currentY)
					}
				} else {
					allWiresMap[mapString] = 1
				}
				currentWireMap[mapString] = true
			}
			currentX += length
			break
		default:
			log.Fatalf("Should not reach here")
		}
	}
	return manhattanDist
}

// Given a current manhattan distance and an intersection point, returns the smaller manhattan distance
func minManhattan(manhattanDist int, currentX int, currentY int) int {
	if manhattanDist == -1 {
		return int(math.Abs(float64(currentX)) + math.Abs(float64(currentY)))
	}
	return int(math.Min(float64(manhattanDist), math.Abs(float64(currentX))+math.Abs(float64(currentY))))

}

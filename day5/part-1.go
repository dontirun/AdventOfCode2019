package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

//Prints the solution for Day 5 Part 1
func partOne() {
	input := 1
	programValues := parseProgramValues()
	writeTracker := map[int]bool{}
	for position := 0; position < len(programValues); {
		if programValues[position] == 99 {
			break
		}
		position += editValue(position, programValues, input, writeTracker)
	}
}

func parseProgramValues() []int {
	file, err := os.Open("./inputs/d5-p1-p2")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	fileCsv, err := ioutil.ReadAll(file)
	fileCsvString := string(fileCsv)
	stringValues := strings.Split(fileCsvString, ",")
	intValues := []int{}
	for _, val := range stringValues {
		intValue, _ := strconv.Atoi(val)
		intValues = append(intValues, intValue)
	}
	return intValues
}

// Edits the slice containing the program values
func editValue(position int, pv []int, input int, writeTracker map[int]bool) int {
	command, modes := discoverCommand(pv[position])
	switch {
	case command == 1:
		val1 := pv[valueSelector(position+1, pv, modes[2:], writeTracker)]
		val2 := pv[valueSelector(position+2, pv, modes[1:2], writeTracker)]
		pv[valueSelector(position+3, pv, modes[:1], writeTracker)] = val1 + val2
		return 4
	case command == 2:
		val1 := pv[valueSelector(position+1, pv, modes[2:], writeTracker)]
		val2 := pv[valueSelector(position+2, pv, modes[1:2], writeTracker)]
		pv[valueSelector(position+3, pv, modes[:1], writeTracker)] = val1 * val2
		return 4
	case command == 3:
		pv[valueSelector(position+1, pv, modes[0:], writeTracker)] = input
		return 2
	case command == 4:
		fmt.Println(pv[valueSelector(position+1, pv, modes[0:], writeTracker)])
		return 2
	default:
		fmt.Println("Edit Value: Shouldn't be reaching here")
	}
	return -1
}

func valueSelector(position int, pv []int, mode string, writeTracker map[int]bool) int {
	_, ok := writeTracker[position]
	switch {
	case ok:
		which := pv[position]
		return which
	case mode == "0":
		which := pv[position]
		return which
	case mode == "1":
		which := position
		return which
	}
	fmt.Println("Value Selector: Shouldn't be reaching here")
	return -1
}
func discoverCommand(fullCommand int) (command int, modes string) {
	command = fullCommand % 10
	switch {
	case command == 1:
		modes = discoverModes(fullCommand, 4)
	case command == 2:
		modes = discoverModes(fullCommand, 4)
	case command == 3:
		modes = discoverModes(fullCommand, 2)
	case command == 4:
		modes = discoverModes(fullCommand, 2)
	default:
		fmt.Println("Discover Command: Shouldn't be reaching here")
	}
	return
}

func discoverModes(fullCommand int, length int) string {
	modeCheck := int(math.Pow10(length + 1))
	if fullCommand < modeCheck {
		fullCommand = fullCommand + modeCheck
	}
	fullCommandString := strconv.Itoa(fullCommand)
	return fullCommandString[1 : len(fullCommandString)-2]
}

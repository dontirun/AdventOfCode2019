package day5

import (
	"fmt"
)

//Prints the solution for Day 5 Part 2
func partTwo() {
	input := 5
	programValues := parseProgramValues()
	writeTracker := map[int]bool{}
	for position := 0; position < len(programValues); {
		if programValues[position] == 99 {
			break
		}
		position += editValue2(position, programValues, input, writeTracker)
	}
}

// Edits the slice containing the program values
func editValue2(position int, pv []int, input int, writeTracker map[int]bool) int {
	command, modes := discoverCommand2(pv[position])
	switch {
	case command == 1:
		val1 := pv[valueSelector(position+1, pv, modes[2:], writeTracker)]
		val2 := pv[valueSelector(position+2, pv, modes[1:2], writeTracker)]
		selectedVal := valueSelector(position+3, pv, modes[:1], writeTracker)
		writeTracker[selectedVal] = true
		pv[selectedVal] = val1 + val2
		return 4
	case command == 2:
		val1 := pv[valueSelector(position+1, pv, modes[2:], writeTracker)]
		val2 := pv[valueSelector(position+2, pv, modes[1:2], writeTracker)]
		selectedVal := valueSelector(position+3, pv, modes[:1], writeTracker)
		writeTracker[selectedVal] = true
		pv[selectedVal] = val1 * val2
		return 4
	case command == 3:
		selectedVal := valueSelector(position+1, pv, modes[0:], writeTracker)
		writeTracker[selectedVal] = true
		pv[selectedVal] = input
		return 2
	case command == 4:
		fmt.Println(pv[valueSelector(position+1, pv, modes[0:], writeTracker)])
		return 2
	case command == 5:
		valTruth := pv[valueSelector(position+1, pv, modes[1:], writeTracker)]
		if valTruth != 0 {
			jumpVal := pv[valueSelector(position+2, pv, modes[0:1], writeTracker)]
			return jumpVal - position
		}
		return 3
	case command == 6:
		valTruth := pv[valueSelector(position+1, pv, modes[1:], writeTracker)]
		if valTruth == 0 {
			jumpVal := pv[valueSelector(position+2, pv, modes[0:1], writeTracker)]
			return jumpVal - position
		}
		return 3
	case command == 7:
		val1 := pv[valueSelector(position+1, pv, modes[2:], writeTracker)]
		val2 := pv[valueSelector(position+2, pv, modes[1:2], writeTracker)]
		selectedVal := valueSelector(position+3, pv, modes[:1], writeTracker)
		writeTracker[selectedVal] = true
		if val1 < val2 {
			pv[selectedVal] = 1
		} else {
			pv[selectedVal] = 0
		}
		return 4
	case command == 8:
		val1 := pv[valueSelector(position+1, pv, modes[2:], writeTracker)]
		val2 := pv[valueSelector(position+2, pv, modes[1:2], writeTracker)]
		selectedVal := valueSelector(position+3, pv, modes[:1], writeTracker)
		writeTracker[selectedVal] = true
		if val1 == val2 {
			pv[selectedVal] = 1
		} else {
			pv[selectedVal] = 0
		}
		return 4
	default:
		fmt.Println("Edit Value 2: Shouldn't be reaching here")
		return -1
	}
}

func discoverCommand2(fullCommand int) (command int, modes string) {
	command = fullCommand % 10
	switch {
	case command == 1:
		modes = discoverModes(fullCommand, 4)
		break
	case command == 2:
		modes = discoverModes(fullCommand, 4)
		break
	case command == 3:
		modes = discoverModes(fullCommand, 2)
		break
	case command == 4:
		modes = discoverModes(fullCommand, 2)
		break
	case command == 5:
		modes = discoverModes(fullCommand, 3)
		break
	case command == 6:
		modes = discoverModes(fullCommand, 3)
		break
	case command == 7:
		modes = discoverModes(fullCommand, 4)
		break
	case command == 8:
		modes = discoverModes(fullCommand, 4)
		break
	default:
		fmt.Println("Discover Command 2: Shouldn't be reaching here")
	}
	return
}

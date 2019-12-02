package day2

import (
	"fmt"
)

var nounMax = 99
var verbMax = 99
var desiredOutput = 19690720

//Prints the solution for Day 2 Part 2
func partTwo() {
	programValues := parseProgramValues()

	for noun := 0; noun <= nounMax; noun++ {
		for verb := 0; verb <= verbMax; verb++ {
			if checkDesiredOutput(programValues, noun, verb) {
				fmt.Printf("%d \n", 100*noun+verb)
				return
			}
		}
	}
}

//Checks if the input noun and verb generate the desired output
func checkDesiredOutput(programValues []int, noun int, verb int) (isDesiredOutput bool) {
	currentProgramValues := make([]int, len(programValues))
	copy(currentProgramValues, programValues)
	currentProgramValues[1] = noun
	currentProgramValues[2] = verb
	for position := 0; position < len(currentProgramValues); position += 4 {
		if currentProgramValues[position] == 99 {
			break
		}
		editValue(position, &currentProgramValues)
	}
	if currentProgramValues[0] == desiredOutput {
		return true
	}
	return false
}

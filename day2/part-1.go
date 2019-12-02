package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

//Prints the solution for Day 2 Part 1
func partOne() {
	programValues := parseProgramValues()
	//modify input for problem 1
	programValues[1] = 12
	programValues[2] = 2
	for position := 0; position < len(programValues); position += 4 {
		if programValues[position] == 99 {
			break
		}
		editValue(position, &programValues)
	}
	fmt.Printf("%v \n", programValues[0])
}

func parseProgramValues() (programValues []int) {
	file, err := os.Open("./inputs/d2-p1-p2")
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	fileCsv, err := ioutil.ReadAll(file)
	fileCsvString := string(fileCsv)
	programValuesString := strings.Split(fileCsvString, ",")
	for _, val := range programValuesString {
		value, err := strconv.Atoi(val)
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		programValues = append(programValues, value)
	}
	file.Close()
	return
}

// Edits the slice containing the program values
func editValue(position int, programValues *[]int) {
	pv := *programValues
	switch {
	case pv[position] == 1:
		pv[pv[position+3]] = pv[pv[position+1]] + pv[pv[position+2]]
		break
	case pv[position] == 2:
		pv[pv[position+3]] = pv[pv[position+1]] * pv[pv[position+2]]
		break
	default:
		fmt.Println("Shouldn't be reaching here")
	}
}

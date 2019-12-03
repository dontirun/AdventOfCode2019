package main

import (
	"fmt"

	day1 "github.com/dontirun/AdventOfCode2019/day1"
	day2 "github.com/dontirun/AdventOfCode2019/day2"
	day3 "github.com/dontirun/AdventOfCode2019/day3"
)

var days = []func(){day1.Solutions, day2.Solutions, day3.Solutions}

func main() {
	for day, solution := range days {
		fmt.Printf("Details for Day %d: \n", day+1)
		solution()
	}
}

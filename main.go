package main

import (
	"fmt"

	"github.com/dontirun/AdventOfCode2019/day1"
	"github.com/dontirun/AdventOfCode2019/day2"
	"github.com/dontirun/AdventOfCode2019/day3"
	"github.com/dontirun/AdventOfCode2019/day4"
	"github.com/dontirun/AdventOfCode2019/day5"
)

var days = []func(){day1.Solutions, day2.Solutions, day3.Solutions, day4.Solutions, day5.Solutions}

func main() {
	for day, solution := range days {
		fmt.Printf("\n Details for Day %d: \n", day+1)
		solution()
	}
}

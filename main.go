package main

import (
	"os"
	"strconv"

	Days "github.com/broemp/adventofcode2022/Days"
)

func main() {
	args := os.Args[1:]

	// if len(args) == 0 || len(args) > 1 {
	// 	println("Please provide a single Day!")
	// 	return
	// }

	for _, dayString := range args {

		day, err := strconv.Atoi(dayString)
		if err != nil {
			println("Please provide an int as Day!")
		}
		if day < 10 {
			dayString = "0" + dayString
		}
		println("Day " + dayString)

		switch day {
		case 1:
			Days.Day01Main()
		case 2:
			Days.Day02Main()
		case 3:
			Days.Day03Main()
		default:
			println("Day not implemented yet!")
		}

		println("----------------------")
	}
}

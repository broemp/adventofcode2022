package Days

import (
	"fmt"
	"log"
	"strconv"

	Helper "github.com/broemp/adventofcode2022/Helper"
)

type elf struct {
	id             int
	snacks         []int
	totallCalories int
}

func newElf(id int) *elf {
	e := elf{id: id}
	return &e
}

func (e elf) totallCaloriesCalc() int {
	for _, calories := range e.snacks {
		e.totallCalories += calories
	}
	return e.totallCalories
}

func Day01Main() {
	dataPath := "./Data/day01.txt"
	data := Helper.ParseFile(dataPath)

	counter := 0

	currentElf := newElf(counter)
	var elfList []*elf

	for _, line := range data {

		if line == "" {
			counter++
			currentElf = newElf(counter)
			elfList = append(elfList, currentElf)
		} else {

			currentCalories, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			currentElf.snacks = append(currentElf.snacks, currentCalories)
		}
	}

	var m int
	for i, cuEl := range elfList {
		currentCal := cuEl.totallCaloriesCalc()
		if i == 0 || currentCal > m {
			m = currentCal
		}
	}
	println("Max Calories: ", m)
	var topList [3]int
	for i, cuEl := range elfList {
		currentCal := cuEl.totallCaloriesCalc()
		if i <= 3 || currentCal > topList[2] {

			if currentCal > topList[0] {
				tempFirst := topList[0]
				topList[0] = currentCal
				tempSecond := topList[1]
				topList[1] = tempFirst
				topList[2] = tempSecond
			} else if currentCal > topList[1] {
				tempSecond := topList[1]
				topList[1] = currentCal
				topList[2] = tempSecond
			} else {
				topList[2] = currentCal
			}

		}
	}
	println("Top Three Total: ", topList[0]+topList[1]+topList[2])
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

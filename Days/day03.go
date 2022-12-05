package Days

import (
	Helper "github.com/broemp/adventofcode2022/Helper"
)

func Day03Main() {
	dataPath := "./Data/day03.txt"
	data := Helper.ParseFile(dataPath)
	println(backpackDuplicateChecker(data))
	println(backpackGroupBadgeFinder(data))
}

func backpackDuplicateChecker(data []string) int {
	var result int

	var duplicates []rune
	for _, line := range data {

		var firstComp, secondComp []rune

		for i, item := range line {

			if i < len(line)/2 {
				firstComp = append(firstComp, item)
			} else {
				secondComp = append(secondComp, item)
			}
		}

		breaker := false
		for _, firstLetter := range firstComp {
			for _, secondLetter := range secondComp {
				if firstLetter == secondLetter {
					duplicates = append(duplicates, firstLetter)
					breaker = true
					break
				}
			}
			if breaker {
				break
			}
		}
	}
	result = calculatePoints(duplicates)
	return result
}

func backpackGroupBadgeFinder(data []string) (result int) {

	var dupesResult []rune
	var firstLine string
	var dupes []rune
	for lineNumber, line := range data {

		if lineNumber%3 == 0 {
			dupes = nil
			firstLine = line
		} else if lineNumber%3 == 1 {
			for _, item := range line {
				for _, item2 := range firstLine {
					if item == item2 {
						dupes = append(dupes, item)
						break
					}
				}

			}
		} else {
			breaker := false
			for _, item := range line {
				for _, item2 := range dupes {
					if item == item2 {
						dupesResult = append(dupesResult, item)
						breaker = true
						break
					}
				}
				if breaker {
					break
				}
			}
		}

	}
	result = calculatePoints(dupesResult)
	return
}

func calculatePoints(duplicates []rune) (result int) {
	for _, num := range duplicates {
		if num <= 96 {
			result += (int(num) - 38)
		} else {
			result += (int(num) - 96)
		}
	}
	return
}

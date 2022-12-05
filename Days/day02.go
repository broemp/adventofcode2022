package Days

import (
	"strings"

	Helper "github.com/broemp/adventofcode2022/Helper"
)

const DRAW = 3
const WIN = 6

func Day02Main() {
	dataPath := "./Data/day02.txt"
	data := Helper.ParseFile(dataPath)
	println("First Round: ", winningCounter(data, 1))
	println("Second Round: ", winningCounter(data, 2))
}

func winningCounter(data []string, tactic int8) (winnings int) {

	for _, line := range data {
		var game [2]int
		switch tactic {
		case 1:
			game = gameParserTacticOne(line)
		case 2:
			game = gameParserTacticTwo(line)
		}

		result := game[0] - game[1]

		if result == 0 {
			winnings = winnings + DRAW + game[1]
		} else if result == -1 || result == 2 {
			winnings = winnings + WIN + game[1]
		} else if result == 1 || result == -2 {
			winnings = winnings + game[1]
		}
	}
	return
}

func gameParserTacticOne(line string) [2]int {
	game := strings.Fields(line)
	var gameNum [2]int

	switch game[0] {
	case "A":
		gameNum[0] = 1
	case "B":
		gameNum[0] = 2
	case "C":
		gameNum[0] = 3
	}

	switch game[1] {
	case "X":
		gameNum[1] = 1
	case "Y":
		gameNum[1] = 2
	case "Z":
		gameNum[1] = 3
	}
	return gameNum
}

func gameParserTacticTwo(line string) [2]int {
	game := strings.Fields(line)
	var gameNum [2]int

	switch game[0] {
	case "A":
		gameNum[0] = 1
	case "B":
		gameNum[0] = 2
	case "C":
		gameNum[0] = 3
	}

	switch game[1] {
	case "X":
		gameNum[1] = (gameNum[0] + 2) % 4
		if gameNum[1] <= 1 {
			gameNum[1]++
		}
	case "Y":
		gameNum[1] = gameNum[0]
	case "Z":
		gameNum[1] = (gameNum[0] + 1) % 4
		if gameNum[1] == 0 {
			gameNum[1]++
		}
	}

	return gameNum
}

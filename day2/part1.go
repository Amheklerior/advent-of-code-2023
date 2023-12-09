package day2

import (
	"amheklerior.com/advent-of-code-2023/utils"
)

const MAX_RED int = 12
const MAX_GREEN int = 13
const MAX_BLUE int = 14

func areCountsWithinRange(counts []string, maxCount int) bool {
	for _, v := range counts {
		if utils.ToInt(v) > maxCount {
			return false
		}
	}
	return true
}

func isValid(game string) bool {
	// Extract counts for all colored cubes
	reds := getCountsFor(game, "red")
	greens := getCountsFor(game, "green")
	blues := getCountsFor(game, "blue")

	// Invalidate what would be an impossible game
	// based on the MAX_* constraints
	isValidGame := areCountsWithinRange(reds, MAX_RED) &&
		areCountsWithinRange(greens, MAX_GREEN) &&
		areCountsWithinRange(blues, MAX_BLUE)

	return isValidGame
}

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	scanner := utils.Scanner(content)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Extract the game id
		_, prefix := utils.ExtractPrefix(line, `Game\s+\d+:`)
		gameId := utils.GetOccurrence(prefix, `\d+`)

		isValidGame := isValid(line)
		if isValidGame {
			sum += utils.ToInt(gameId)
		}
	}
	return sum
}

func TestP1() {
	utils.Run(2, 1, 8, SolutionPart1)
}

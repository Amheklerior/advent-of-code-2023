package day7

import (
	"amheklerior.com/advent-of-code-2023/day7/game"
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	game := game.NewCamelGame(content, true)
	return game.Score()
}

func TestP2() {
	utils.Run(7, 2, 5905, SolutionPart2)
}

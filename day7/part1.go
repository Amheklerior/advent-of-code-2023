package day7

import (
	"amheklerior.com/advent-of-code-2023/day7/game"
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	game := game.NewCamelGame(content)
	return game.Score()
}

func TestP1() {
	utils.Run(7, 1, 6440, SolutionPart1)
}

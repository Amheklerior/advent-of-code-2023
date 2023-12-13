package day10

import (
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)

	return len(content)
}

func TestP1() {
	utils.Run(10, 1, -1, SolutionPart1)
}

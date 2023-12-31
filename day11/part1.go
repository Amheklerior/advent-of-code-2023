package day11

import (
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	univ := NewUniverse(content)
	return univ.Solve(1)
}

func TestP1() {
	utils.Run(11, 1, 374, SolutionPart1)
}

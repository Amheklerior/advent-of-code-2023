package day11

import (
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	univ := NewUniverse(content)
	return univ.Solve(1000000)
}

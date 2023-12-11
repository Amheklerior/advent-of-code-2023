package day8

import (
	"amheklerior.com/advent-of-code-2023/day8/navigator"
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	instructions := navigator.NewInstructionPath(content)
	network, _ := navigator.NewNetwork(content)
	return network.Traverse(instructions)
}

func TestP1() {
	utils.Run(8, 1, 2, SolutionPart1)
}

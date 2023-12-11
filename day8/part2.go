package day8

import (
	"amheklerior.com/advent-of-code-2023/day8/navigator"
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	instructions := navigator.NewInstructionPath(content)
	network, entryPoints := navigator.NewNetwork(content)
	var paths []int
	for _, entrypoint := range entryPoints {
		hops := network.FindPathLenghtFrom(entrypoint, instructions)
		paths = append(paths, hops)
	}
	return (utils.IntSlice(paths)).LeastCommonMultiple()
}

func TestP2() {
	utils.Run(8, 2, 6, SolutionPart2)
}

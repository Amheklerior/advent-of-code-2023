package day10

import (
	"amheklerior.com/advent-of-code-2023/day10/maze"
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	terrain := maze.BuildTerrain(content)
	pipePath := terrain.BuildPipePath()
	return len(pipePath) / 2
}

func TestP1() {
	utils.Run(10, 1, 8, SolutionPart1)
}

package day10

import (
	"amheklerior.com/advent-of-code-2023/day10/maze"
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	terrain := maze.BuildTerrain(content)
	pipeLoop := terrain.BuildPipeLoop()
	fsm := maze.NewFSM(terrain, pipeLoop)
	return fsm.Solve()
}

func TestP2() {
	utils.Run(10, 2, 10, SolutionPart2)
}

package day10

import (
	"fmt"

	"amheklerior.com/advent-of-code-2023/day10/maze"
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	terrain := maze.BuildTerrain(content)
	fmt.Printf("%v", terrain.String())
	pipeLoop := maze.BuildPipeLoop(&terrain)
	terrain.CleanupFromJunkPipes(pipeLoop)
	fmt.Printf("%v", terrain.String())
	fsm := maze.NewFSM(terrain, pipeLoop)
	return fsm.Solve()
}

func TestP2() {
	utils.Run(10, 2, 10, SolutionPart2)
}

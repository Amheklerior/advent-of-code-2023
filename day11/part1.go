package day11

import (
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	space := NewSpace(content)
	space = space.Expand()
	galaxies := space.GalaxiesMap()
	sum := 0
	for i, galaxy := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += Distance(galaxy, galaxies[j])
		}
	}
	return sum
}

func TestP1() {
	utils.Run(11, 1, 374, SolutionPart1)
}

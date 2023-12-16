package day11

import (
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	univ := NewUniverse(content)
	galaxies := univ.GetGalaxies()

	sum := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += univ.DistanceBetween(galaxies[i], galaxies[j], 1)
		}
	}
	return sum
}

func TestP1() {
	utils.Run(11, 1, 374, SolutionPart1)
}

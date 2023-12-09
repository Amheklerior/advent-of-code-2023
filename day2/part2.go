package day2

import (
	"sort"

	"amheklerior.com/advent-of-code-2023/utils"
)

func findMax(counts []string) int {
	sort.Slice(counts, func(i, j int) bool {
		return utils.ToInt(counts[i]) < utils.ToInt(counts[j])
	})
	max := utils.ToInt(counts[len(counts)-1])
	return max
}

func fewestNumberOfCubes(game string) (int, int, int) {
	// Extract the counts for all colored cubes
	reds := getCountsFor(game, "red")
	greens := getCountsFor(game, "green")
	blues := getCountsFor(game, "blue")

	// Find the highest count for each colored cube
	maxRed := findMax(reds)
	maxGreen := findMax(greens)
	maxBlue := findMax(blues)

	return maxRed, maxGreen, maxBlue
}

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)

	sum := 0
	scanner := utils.Scanner(content)
	for scanner.Scan() {
		line := scanner.Text()
		red, green, blue := fewestNumberOfCubes(line)
		sum += red * green * blue
	}
	return sum
}

func TestP2() {
	utils.Run(2, 2, 2286, SolutionPart2)
}

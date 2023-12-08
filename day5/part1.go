package day5

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"amheklerior.com/advent-of-code-2023/day5/lib"
	"amheklerior.com/advent-of-code-2023/utils"
)

func extractSeeds(input string) []int {
	var seeds []int

	// grab the first line without the prefix
	scanner := utils.Scanner(input)
	scanner.Scan()
	line := scanner.Text()
	prefix := regexp.MustCompile(`seeds:`).FindString(line)
	line = strings.TrimPrefix(line, prefix)

	// extract all seeds
	ids := strings.Fields(line)
	for _, id := range ids {
		seed, _ := strconv.Atoi(id)
		seeds = append(seeds, seed)
	}
	return seeds
}

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)

	seeds := extractSeeds(content)
	pipeline := lib.BuildDataStructures(content)

	minLocation := math.MaxInt
	for _, seed := range seeds {
		location := lib.PassThroughPipeline(seed, pipeline)
		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func TestP1() {
	utils.Run(5, 1, 35, SolutionPart1)
}

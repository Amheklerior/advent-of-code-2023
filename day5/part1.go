package day5

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"amheklerior.com/advent-of-code-2023/day5/lib"
)

func extractSeeds(line string) []int {
	var seeds []int
	prefix := regexp.MustCompile(`seeds:`).FindString(line)
	line = strings.TrimPrefix(line, prefix)
	ids := strings.Fields(line)
	for _, id := range ids {
		seed, _ := strconv.Atoi(id)
		seeds = append(seeds, seed)
	}
	return seeds
}

func SolutionPart1(path string) int {
	content := lib.ReadFile(path)

	seeds := extractSeeds(strings.SplitN(content, "\n", 1)[0])
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
	fmt.Println("Day 5 / Part 1: Test")
	expected := 35
	result := SolutionPart1("./day5/data/p1-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

package day5

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"amheklerior.com/advent-of-code-2023/day5/lib"
)

func extractSeedsInRanges(input string) []lib.Range {
	var seeds []lib.Range

	// grab first line without the prefix
	scanner := lib.Scanner(input)
	scanner.Scan()
	line := scanner.Text()
	prefix := regexp.MustCompile(`seeds:`).FindString(line)
	line = strings.TrimPrefix(line, prefix)

	// extract all seeds ranges
	pairs := regexp.MustCompile(`\s*\d+\s+\d+\s*`).FindAllString(line, -1)
	for _, pair := range pairs {
		data := strings.Fields(pair)
		rangeStart, _ := strconv.Atoi(data[0])
		rangeLenght, _ := strconv.Atoi(data[1])
		seeds = append(seeds, lib.NewRange(rangeStart, rangeStart+rangeLenght))
	}

	return seeds
}

// computation time: ~6m40s
func SolutionPart2(path string) int {
	content := lib.ReadFile(path)
	seeds := extractSeedsInRanges(content)
	pipeline := lib.BuildDataStructures(content)

	// Process each seed through the pipeline and get the min location
	var wg sync.WaitGroup
	minLocation := math.MaxInt
	for _, seedRange := range seeds {
		for seed := seedRange.Start; seed < seedRange.End; seed++ {
			wg.Add(1)
			go func(seed int) {
				defer wg.Done()
				location := lib.PassThroughPipeline(seed, pipeline)
				if location < minLocation {
					minLocation = location
				}
			}(seed)
		}
	}

	wg.Wait()
	return minLocation
}

func TestP2() {
	fmt.Println("Day 5 / Part 2: Test")
	expected := 46
	result := SolutionPart2("./day5/data/p2-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

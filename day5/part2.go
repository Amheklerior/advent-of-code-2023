package day5

import (
	"math"
	"strings"
	"sync"

	"amheklerior.com/advent-of-code-2023/day5/lib"
	"amheklerior.com/advent-of-code-2023/utils"
)

func extractSeedsInRanges(input string) []lib.Range {
	var seeds []lib.Range

	// grab first line without the prefix
	scanner := utils.Scanner(input)
	scanner.Scan()
	line := scanner.Text()
	line, _ = utils.ExtractPrefix(line, `seeds:`)

	// extract all seeds ranges
	pairs := utils.GetOccurrences(line, `\s*\d+\s+\d+\s*`)
	for _, pair := range pairs {
		data := strings.Fields(pair)
		rangeStart := utils.ToInt(data[0])
		rangeLenght := utils.ToInt(data[1])
		seeds = append(seeds, lib.NewRange(rangeStart, rangeStart+rangeLenght))
	}

	return seeds
}

// computation time: ~6m40s
func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
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
	utils.Run(5, 2, 46, SolutionPart2)
}

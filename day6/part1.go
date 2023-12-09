package day6

import (
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

func extractRaces(content string) []Race {
	var races []Race
	var times []string
	var distances []string

	scanner := utils.Scanner(content)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Time") {
			line, _ = utils.ExtractPrefix(line, `Time:`)
			times = append(times, strings.Fields(line)...)
			continue
		}

		if strings.Contains(line, "Distance") {
			line, _ = utils.ExtractPrefix(line, `Distance:`)
			distances = append(distances, strings.Fields(line)...)
			continue
		}
	}

	for i, time := range times {
		races = append(races, Race{
			utils.ToInt(time),
			utils.ToInt(distances[i]),
		})
	}

	return races
}

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	races := extractRaces(content)

	result := 1
	for _, race := range races {
		waysToWin := getWaysToWin(race)
		result *= waysToWin
	}
	return result
}

func TestP1() {
	utils.Run(6, 1, 288, SolutionPart1)
}

package day6

import (
	"regexp"
	"strconv"
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

func extractRace(content string) Race {
	var time int
	var distance int

	scanner := utils.Scanner(content)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Time") {
			prefix := regexp.MustCompile(`Time:`).FindString(line)
			line = strings.TrimPrefix(line, prefix)
			partials := strings.Fields(line)
			time, _ = strconv.Atoi(strings.Join(partials, ""))
			continue
		}

		if strings.Contains(line, "Distance") {
			prefix := regexp.MustCompile(`Distance:`).FindString(line)
			line = strings.TrimPrefix(line, prefix)
			partials := strings.Fields(line)
			distance, _ = strconv.Atoi(strings.Join(partials, ""))
			continue
		}
	}

	return Race{time, distance}
}

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	race := extractRace(content)
	return getWaysToWin(race)
}

func TestP2() {
	utils.Run(6, 2, 71503, SolutionPart2)
}

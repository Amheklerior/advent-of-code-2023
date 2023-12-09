package day6

import (
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
			line, _ = utils.ExtractPrefix(line, `Time:`)
			partials := strings.Fields(line)
			time = utils.ToInt(strings.Join(partials, ""))
			continue
		}

		if strings.Contains(line, "Distance") {
			line, _ = utils.ExtractPrefix(line, `Distance:`)
			partials := strings.Fields(line)
			distance = utils.ToInt(strings.Join(partials, ""))
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

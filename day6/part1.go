package day6

import (
	"regexp"
	"strconv"
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

func simulateDistance(totTime, chargeTime int) int {
	const VELOCITY_CHARGE_PER_MS = 1
	v := VELOCITY_CHARGE_PER_MS * chargeTime
	d := v * (totTime - chargeTime)
	return d
}

func findLowestChargeTime(totTime, record int) int {
	chargeTime := 1
	distance := simulateDistance(totTime, chargeTime)

	for distance <= record {
		chargeTime++
		distance = simulateDistance(totTime, chargeTime)
	}

	return chargeTime
}

func findHighestChargeTime(totTime, record int) int {
	chargeTime := totTime - 1
	distance := simulateDistance(totTime, chargeTime)

	for distance <= record {
		chargeTime--
		distance = simulateDistance(totTime, chargeTime)
	}

	return chargeTime
}

type Race struct {
	totTime, record int
}

func extractData(content string) []Race {
	var races []Race
	var times []string
	var distances []string

	scanner := utils.Scanner(content)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "Time") {
			prefix := regexp.MustCompile(`Time:`).FindString(line)
			line = strings.TrimPrefix(line, prefix)
			times = append(times, strings.Fields(line)...)
			continue
		}

		if strings.Contains(line, "Distance") {
			prefix := regexp.MustCompile(`Distance:`).FindString(line)
			line = strings.TrimPrefix(line, prefix)
			distances = append(distances, strings.Fields(line)...)
			continue
		}
	}

	for i, t := range times {
		time, _ := strconv.Atoi(t)
		distance, _ := strconv.Atoi(distances[i])
		races = append(races, Race{time, distance})
	}

	return races
}

func getWaysToWin(race Race) int {
	return findHighestChargeTime(race.totTime, race.record) - findLowestChargeTime(race.totTime, race.record) + 1
}

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	races := extractData(content)

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

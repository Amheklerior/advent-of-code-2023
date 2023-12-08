package day4

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

func cardScore(card string) int {
	// remove the prefix
	prefix := regexp.MustCompile(`Card \d+: `).FindString(card)
	card = strings.TrimPrefix(card, prefix)

	// split the number lists
	list := strings.Split(card, "|")
	winnings := strings.Fields(strings.Trim(list[0], " "))
	hand := strings.Fields(strings.Trim(list[1], " "))

	// count winning numbers
	wins := 0
	for _, num := range hand {
		isWinning := slices.ContainsFunc(winnings, func(el string) bool { return el == num })
		if isWinning {
			wins++
		}
	}

	// return the totaling score
	if wins == 0 || wins == 1 {
		return wins
	}
	return int(math.Pow(2, float64(wins-1)))
}

func SolutionPart1(path string) int {
	f, e := os.Open(path)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += cardScore(line)
	}
	return sum
}

func TestP1() {
	utils.Run(4, 1, 13, SolutionPart1)
}

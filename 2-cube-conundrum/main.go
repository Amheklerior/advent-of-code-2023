package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const PATH string = "./input.txt"
const MAX_RED int = 12
const MAX_GREEN int = 13
const MAX_BLUE int = 14

func getCountsFor(str, color string) []string {
	digitsRegex := regexp.MustCompile(`\d+`)
	r := regexp.MustCompile(fmt.Sprintf(`\d+ %s`, color))
	return digitsRegex.FindAllString(strings.Join(r.FindAllString(str, -1), " "), -1)
}

func areCountsWithinRange(counts []string, maxCount int) bool {
	for _, v := range counts {
		count, _ := strconv.Atoi(v)
		if count > maxCount {
			return false
		}
	}
	return true
}

func main() {
	f, e := os.Open(PATH)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Extract the game id and counts for all colored cubes
		gameId, _ := strings.CutPrefix(strings.Split(line, ":")[0], "Game ")
		id, _ := strconv.Atoi(gameId)
		reds := getCountsFor(line, "red")
		greens := getCountsFor(line, "green")
		blues := getCountsFor(line, "blue")

		// Invalidate what would be an impossible game
		// based on the MAX_* constraints
		isValidGame := areCountsWithinRange(reds, MAX_RED) &&
			areCountsWithinRange(greens, MAX_GREEN) &&
			areCountsWithinRange(blues, MAX_BLUE)

		// Sum up valid game ids
		if isValidGame {
			sum += id
		}
	}

	fmt.Println(sum)
}

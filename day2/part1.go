package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const MAX_RED int = 12
const MAX_GREEN int = 13
const MAX_BLUE int = 14

func areCountsWithinRange(counts []string, maxCount int) bool {
	for _, v := range counts {
		count, _ := strconv.Atoi(v)
		if count > maxCount {
			return false
		}
	}
	return true
}

func isValid(game string) bool {
	// Extract counts for all colored cubes
	reds := getCountsFor(game, "red")
	greens := getCountsFor(game, "green")
	blues := getCountsFor(game, "blue")

	// Invalidate what would be an impossible game
	// based on the MAX_* constraints
	isValidGame := areCountsWithinRange(reds, MAX_RED) &&
		areCountsWithinRange(greens, MAX_GREEN) &&
		areCountsWithinRange(blues, MAX_BLUE)

	return isValidGame
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

		// Extract the game id
		gameId, _ := strings.CutPrefix(strings.Split(line, ":")[0], "Game ")
		id, _ := strconv.Atoi(gameId)

		isValidGame := isValid(line)
		if isValidGame {
			sum += id
		}
	}
	return sum
}

func TestP1() {
	fmt.Println("Day 2 / Part 1: Test")
	expected := 8
	result := SolutionPart1("./day2/data/p1-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

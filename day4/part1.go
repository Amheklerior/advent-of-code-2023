package day4

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strings"
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
	fmt.Println("Day 4 / Part 1: Test")
	expected := 13
	result := SolutionPart1("./day4/data/p1-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

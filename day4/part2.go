package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id      int
	content string
	count   int
}

func (card *Card) winningNumCount() int {
	// split the number lists
	list := strings.Split(card.content, "|")
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

	return wins
}

func buildDeck(f *os.File) []*Card {
	var cards []*Card

	// occupy slot 0 to keep alignment with card id
	cards = append(cards, &Card{0, "", 0})

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		prefix := regexp.MustCompile(`Card\s+\d+:\s+`).FindString(line)
		id, _ := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(prefix))
		line = strings.TrimPrefix(line, prefix)
		cards = append(cards, &Card{id, line, 1})
	}

	return cards
}

func SolutionPart2(path string) int {
	f, e := os.Open(path)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	deck := buildDeck(f)
	maxId := len(deck)

	sum := len(deck) - 1

	// for each card in the deck...
	for id, card := range deck {
		if id == 0 {
			continue // skip first slot (no card)
		}

		wins := card.winningNumCount()

		// count the won cards and add them back in the deck
		for i := 1; i <= wins; i++ {
			if card.id+i > maxId {
				break
			}
			deck[id+i].count += card.count
			sum += card.count
		}
	}

	return sum
}

func TestP2() {
	fmt.Println("Day 4 / Part 2: Test")
	expected := 30
	result := SolutionPart2("./day4/data/p2-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

package day4

import (
	"slices"
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
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

func buildDeck(input string) []*Card {
	var cards []*Card

	// occupy slot 0 to keep alignment with card id
	cards = append(cards, &Card{0, "", 0})

	scanner := utils.Scanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		line, prefix := utils.ExtractPrefix(line, `Card\s+\d+:\s+`)
		id := utils.ToInt(utils.GetOccurrence(prefix, `\d+`))
		cards = append(cards, &Card{id, line, 1})
	}

	return cards
}

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)

	deck := buildDeck(content)
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
	utils.Run(4, 2, 30, SolutionPart2)
}

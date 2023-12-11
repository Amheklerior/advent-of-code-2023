package game

import (
	"slices"
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

type CamelGame struct {
	hands      []Hand
	useJollies bool
}

func NewCamelGame(input string, useJollies bool) CamelGame {
	var hands []Hand
	if useJollies {
		input = strings.ReplaceAll(input, string(JACK), string(JOLLY))
	}
	scanner := utils.Scanner(input)
	for scanner.Scan() {
		hands = append(hands, NewHand(scanner.Text()))
	}
	return CamelGame{hands, useJollies}
}

func (game *CamelGame) Score() int {
	sum := 0
	rankHands(game.hands)
	for rank, hand := range game.hands {
		sum += hand.Bid * (rank + 1)
	}
	return sum
}

func rankHands(hands []Hand) []Hand {
	slices.SortStableFunc(hands, func(h1, h2 Hand) int {
		if h1.Type == h2.Type {
			for i, c1 := range h1.Cards {
				c2 := h2.Cards[i]
				if c1 == c2 {
					continue
				}
				return compareCards(c1, c2)
			}
			return 0
		}
		return compareTypes(h1.Type, h2.Type)
	})
	return hands
}

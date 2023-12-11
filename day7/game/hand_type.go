package game

import (
	"slices"
	"sort"
)

type HandType string

const (
	FIVE_OF_A_KIND  HandType = "Five of a kind"
	FOUR_OF_A_KIND  HandType = "Four of a kind"
	FULL_HOUSE      HandType = "Full house"
	THREE_OF_A_KIND HandType = "Three of a kind"
	TWO_PAIRS       HandType = "Two pairs"
	ONE_PAIR        HandType = "One pair"
	HIGH_CARD       HandType = "High card"
)

var handTypeRanking = map[HandType]int{
	FIVE_OF_A_KIND:  7,
	FOUR_OF_A_KIND:  6,
	FULL_HOUSE:      5,
	THREE_OF_A_KIND: 4,
	TWO_PAIRS:       3,
	ONE_PAIR:        2,
	HIGH_CARD:       1,
}

func getHandType(cards [5]Card) HandType {
	// count cards' occurrences
	countByCard := make(map[Card]int)
	for _, card := range cards {
		countByCard[card]++
	}

	// list the counts in an ordered array
	counts := make([]int, 0, len(countByCard))
	for _, count := range countByCard {
		counts = append(counts, count)
	}
	sort.Ints(counts)

	includeJollies := slices.Contains(cards[:], JOLLY)

	// get hand type from the count list
	switch len(counts) {
	case 1:
		// if there's only one count -> all cards are equal
		return FIVE_OF_A_KIND
	case 2:
		// if there's two counts -> can either be that I have:
		// only one card differing from the rest, or
		// a pair differing from the rest
		if includeJollies {
			return FIVE_OF_A_KIND
		}
		if counts[0] == 4 || counts[1] == 4 {
			return FOUR_OF_A_KIND
		}
		return FULL_HOUSE
	case 3:
		// if there's three -> I can have:
		// two different cards differing from the rest, or
		// two pairs of cards differing from the onw card remaining
		if counts[0] == 3 || counts[1] == 3 || counts[2] == 3 {
			if includeJollies {
				return FOUR_OF_A_KIND
			}
			return THREE_OF_A_KIND
		}
		if includeJollies && countByCard[JOLLY] == 2 {
			return FOUR_OF_A_KIND
		}
		if includeJollies && countByCard[JOLLY] == 1 {
			return FULL_HOUSE
		}
		return TWO_PAIRS
	case 4:
		if includeJollies {
			return THREE_OF_A_KIND
		}
		return ONE_PAIR
	case 5:
		if includeJollies {
			return ONE_PAIR
		}
		return HIGH_CARD
	default:
		return HIGH_CARD
	}
}

func compareTypes(type1, type2 HandType) int {
	return handTypeRanking[type1] - handTypeRanking[type2]
}

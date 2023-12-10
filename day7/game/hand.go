package game

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Hand struct {
	Cards [5]Card
	Type  HandType
	Bid   int
}

func NewHand(input string) Hand {
	parts := strings.Fields(input)

	cardsInput := parts[0]
	if len(cardsInput) != 5 {
		log.Fatalf("%v is not a valid hand. There should be 5 cards, found %v cards instead!", cardsInput, len(cardsInput))
	}

	bidPart := parts[1]
	bid, err := strconv.Atoi(bidPart)
	if err != nil {
		log.Fatalf("%v is not a valid bid. Only integers are allowed!", bidPart)
	}

	var cards [5]Card
	for i, c := range cardsInput {
		cards[i] = Card(c)
	}

	return Hand{
		Cards: cards,
		Type:  getHandType(cards),
		Bid:   bid,
	}
}

func (h *Hand) String() string {
	return fmt.Sprintf(
		"Hand(%s%s%s%s%s / %v / bid: %v)",
		string(h.Cards[0]),
		string(h.Cards[1]),
		string(h.Cards[2]),
		string(h.Cards[3]),
		string(h.Cards[4]),
		string(h.Type),
		h.Bid,
	)
}

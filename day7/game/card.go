package game

type Card rune

const (
	ACE   Card = 'A'
	KING  Card = 'K'
	QUEEN Card = 'Q'
	JACK  Card = 'J'
	TEN   Card = 'T'
	NINE  Card = '9'
	EIGHT Card = '8'
	SEVEN Card = '7'
	SIX   Card = '6'
	FIVE  Card = '5'
	FOUR  Card = '4'
	THREE Card = '3'
	TWO   Card = '2'
)

var cardValues = map[Card]int{
	ACE:   14,
	KING:  13,
	QUEEN: 12,
	JACK:  11,
	TEN:   10,
	NINE:  9,
	EIGHT: 8,
	SEVEN: 7,
	SIX:   6,
	FIVE:  5,
	FOUR:  4,
	THREE: 3,
	TWO:   2,
}

func compareCards(card1, card2 Card) int {
	return cardValues[card1] - cardValues[card2]
}

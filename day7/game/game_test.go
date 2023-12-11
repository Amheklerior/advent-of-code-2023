package game

import (
	"fmt"
	"testing"
)

func TestHandConstructor(t *testing.T) {
	// create an entry for each test with the inputs and the expected outputs
	var testCases = []struct {
		input    string
		expected Hand
	}{
		{"32T3K 765", Hand{
			[5]Card{THREE, TWO, TEN, THREE, KING},
			ONE_PAIR,
			765,
		}},
		{"T55J5 684", Hand{
			[5]Card{TEN, FIVE, FIVE, JACK, FIVE},
			THREE_OF_A_KIND,
			684,
		}},
		{"KTJJT 220", Hand{
			[5]Card{KING, TEN, JACK, JACK, TEN},
			TWO_PAIRS,
			220,
		}},
		{"KTA5J 10", Hand{
			[5]Card{KING, TEN, ACE, FIVE, JACK},
			HIGH_CARD,
			10,
		}},
		{"QQQAA 483", Hand{
			[5]Card{QUEEN, QUEEN, QUEEN, ACE, ACE},
			FULL_HOUSE,
			483,
		}},
	}

	// loop through the table and run the same test for each test case
	for _, tt := range testCases {
		testName := fmt.Sprintf("%v should produce %v", tt.input, tt.expected.String())

		// the t.Run() enables running “subtests"
		t.Run(testName, func(t *testing.T) {
			answer := NewHand(tt.input)
			if answer != tt.expected {
				t.Errorf("Got %v, expected %v (for input: %v)", answer, tt.expected, tt.input)
			}
		})
	}
}

func TestGameConstructor(t *testing.T) {
	input := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"
	expected := CamelGame{[]Hand{
		{
			[5]Card{THREE, TWO, TEN, THREE, KING},
			ONE_PAIR,
			765,
		},
		{
			[5]Card{TEN, FIVE, FIVE, JACK, FIVE},
			THREE_OF_A_KIND,
			684,
		},
		{
			[5]Card{KING, KING, SIX, SEVEN, SEVEN},
			TWO_PAIRS,
			28,
		},
		{
			[5]Card{KING, TEN, JACK, JACK, TEN},
			TWO_PAIRS,
			220,
		},
		{
			[5]Card{QUEEN, QUEEN, QUEEN, JACK, ACE},
			THREE_OF_A_KIND,
			483,
		},
	}, false}

	answer := NewCamelGame(input, false)

	if len(answer.hands) != len(expected.hands) {
		t.Errorf("Error: the game should have %v hands, got %v instead",
			len(expected.hands),
			len(answer.hands))
	}

	for i, hand := range answer.hands {
		if hand != expected.hands[i] {
			t.Errorf("Error on hand #%v: expected %v, got %v instead",
				i,
				expected.hands[i].String(),
				hand.String())
		}
	}
}

func TestGameConstructorWithJollies(t *testing.T) {
	input := "32T3K 765\nT55J5 684\nKK677 28\nKTJJT 220\nQQQJA 483"
	expected := CamelGame{[]Hand{
		{
			[5]Card{THREE, TWO, TEN, THREE, KING},
			ONE_PAIR,
			765,
		},
		{
			[5]Card{TEN, FIVE, FIVE, JOLLY, FIVE},
			FOUR_OF_A_KIND,
			684,
		},
		{
			[5]Card{KING, KING, SIX, SEVEN, SEVEN},
			TWO_PAIRS,
			28,
		},
		{
			[5]Card{KING, TEN, JOLLY, JOLLY, TEN},
			FOUR_OF_A_KIND,
			220,
		},
		{
			[5]Card{QUEEN, QUEEN, QUEEN, JOLLY, ACE},
			FOUR_OF_A_KIND,
			483,
		},
	}, true}

	answer := NewCamelGame(input, true)

	if len(answer.hands) != len(expected.hands) {
		t.Errorf("Error: the game should have %v hands, got %v instead",
			len(expected.hands),
			len(answer.hands))
	}

	for i, hand := range answer.hands {
		if hand != expected.hands[i] {
			t.Errorf("Error on hand #%v: expected %v, got %v instead",
				i,
				expected.hands[i].String(),
				hand.String())
		}
	}
}

func TestRankingHands(t *testing.T) {
	input := []Hand{
		{
			[5]Card{THREE, TWO, TEN, THREE, KING},
			ONE_PAIR,
			765,
		},
		{
			[5]Card{TEN, FIVE, FIVE, JACK, FIVE},
			THREE_OF_A_KIND,
			684,
		},
		{
			[5]Card{KING, KING, SIX, SEVEN, SEVEN},
			TWO_PAIRS,
			28,
		},
		{
			[5]Card{KING, TEN, JACK, JACK, TEN},
			TWO_PAIRS,
			220,
		},
		{
			[5]Card{QUEEN, QUEEN, QUEEN, JACK, ACE},
			THREE_OF_A_KIND,
			483,
		},
	}

	expected := []Hand{
		input[4],
		input[1],
		input[2],
		input[3],
		input[0],
	}

	var answer []Hand
	copy(answer, input)
	rankHands(answer)

	for i, hand := range answer {
		if hand != expected[i] {
			t.Errorf("Error on rank #%v: expected %v, got %v instead",
				i,
				expected[i].String(),
				hand.String())
		}
	}
}

func TestRankingHandsWithJollies(t *testing.T) {
	input := []Hand{
		{
			[5]Card{THREE, TWO, TEN, THREE, KING},
			ONE_PAIR,
			765,
		},
		{
			[5]Card{TEN, FIVE, FIVE, JOLLY, FIVE},
			FOUR_OF_A_KIND,
			684,
		},
		{
			[5]Card{KING, KING, SIX, SEVEN, SEVEN},
			TWO_PAIRS,
			28,
		},
		{
			[5]Card{KING, TEN, JOLLY, JOLLY, TEN},
			FOUR_OF_A_KIND,
			220,
		},
		{
			[5]Card{QUEEN, QUEEN, QUEEN, JOLLY, ACE},
			FOUR_OF_A_KIND,
			483,
		},
	}

	expected := []Hand{
		input[3],
		input[4],
		input[1],
		input[2],
		input[0],
	}

	var answer []Hand
	copy(answer, input)
	rankHands(answer)

	for i, hand := range answer {
		if hand != expected[i] {
			t.Errorf("Error on rank #%v: expected %v, got %v instead",
				i,
				expected[i].String(),
				hand.String())
		}
	}
}

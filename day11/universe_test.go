package day11

import (
	"slices"
	"testing"

	"amheklerior.com/advent-of-code-2023/utils"
)

func TestGetGalaxies(t *testing.T) {
	univ := Universe{
		Space{
			{EMPTY, GALAXY, EMPTY},
			{EMPTY, EMPTY, GALAXY},
			{GALAXY, EMPTY, EMPTY},
		},
		make([]int, 0),
		make([]int, 0),
	}
	expected := []Coordinate{
		{1, 0},
		{2, 1},
		{0, 2},
	}

	found := univ.GetGalaxies()

	if !slices.Equal(found, expected) {
		t.Errorf("Got %v, expected %v", found, expected)
	}
}

func TestDistance(t *testing.T) {
	univ := NewUniverse(utils.ReadFile("./data/p1-input.test.txt"))
	univ.GetGalaxies()

	testCases := []struct {
		name            string
		galaxy1         Coordinate
		galaxy2         Coordinate
		expansionFactor int
		expected        int
	}{
		{"Galaxies 5 and 9 are 9 light years apart from each others", Coordinate{1, 5}, Coordinate{4, 9}, 1, 9},
		{"Galaxies 1 and 7 are 15 light years apart from each others", Coordinate{3, 0}, Coordinate{7, 8}, 1, 15},
		{"Galaxies 3 and 6 are 17 light years apart from each others", Coordinate{9, 6}, Coordinate{0, 2}, 1, 17},
		{"Galaxies 8 and 9 are 5 light years apart from each others", Coordinate{0, 9}, Coordinate{4, 9}, 1, 5},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			distance := univ.DistanceBetween(tt.galaxy1, tt.galaxy2, tt.expansionFactor)
			if distance != tt.expected {
				t.Errorf(
					"Got %v, expected %v. For galaxies %v and %v, with expansion factor of %v",
					distance,
					tt.expected,
					tt.galaxy1,
					tt.galaxy2,
					tt.expansionFactor,
				)
			}
		})
	}

}

func TestSolve(t *testing.T) {
	univ := NewUniverse(utils.ReadFile("./data/p1-input.test.txt"))
	univ.GetGalaxies()

	testCases := []struct {
		name            string
		expansionFactor int
		expected        int
	}{
		{"Expansion factor: 1", 1, 374},
		{"Expansion factor: 10", 10, 1030},
		{"Expansion factor: 100", 100, 8410},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			solution := univ.Solve(tt.expansionFactor)
			if solution != tt.expected {
				t.Errorf(
					"Got %v, expected %v. For expansion factor %v",
					solution,
					tt.expected,
					tt.expansionFactor,
				)
			}
		})
	}
}

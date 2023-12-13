package maze

import "testing"

func TestDirVector(t *testing.T) {
	var testCases = []struct {
		testName string
		dir      Direction
		expected Vector
	}{
		{"North points upward", NORTH, Vector{-1, 0}},
		{"South points downward", SOUTH, Vector{1, 0}},
		{"East points to the right", EAST, Vector{0, 1}},
		{"West points to the left", WEST, Vector{0, -1}},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			answer := vector[tt.dir]
			if answer != tt.expected {
				t.Errorf("Got %v, expected %v (for dir: %v)", answer, tt.expected, string(tt.dir))
			}
		})
	}
}

func TestConnections(t *testing.T) {
	var testCases = []struct {
		testName string
		dirFrom  Direction
		dirTo    Direction
		expected bool
	}{
		{"North connects to South", NORTH, SOUTH, true},
		{"North doesn't connect to North", NORTH, NORTH, false},
		{"North doesn't connect to East", NORTH, EAST, false},
		{"North doesn't connect to West", NORTH, WEST, false},
		{"South connects to North", SOUTH, NORTH, true},
		{"South doesn't connect to South", SOUTH, SOUTH, false},
		{"South doesn't connect to East", SOUTH, EAST, false},
		{"South doesn't connect to West", SOUTH, WEST, false},
		{"East connects to West", EAST, WEST, true},
		{"East doesn't connect to East", EAST, EAST, false},
		{"East doesn't connect to North", EAST, NORTH, false},
		{"East doesn't connect to South", EAST, SOUTH, false},
		{"West connects to East", WEST, EAST, true},
		{"West doesn't connect to West", WEST, WEST, false},
		{"West doesn't connect to North", WEST, NORTH, false},
		{"West doesn't connect to South", WEST, SOUTH, false},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			answer := tt.dirFrom.connectsTo(tt.dirTo)
			if answer != tt.expected {
				t.Errorf("Got %v, expected %v (for dir: %v)", answer, tt.expected, string(tt.dirFrom))
			}
		})
	}
}

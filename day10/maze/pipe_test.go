package maze

import (
	"slices"
	"testing"
)

func TestPipeConnectionPoints(t *testing.T) {
	var testCases = []struct {
		testName string
		pipe     Pipe
		expected []Direction
	}{
		{"The vertical pipe has north and south connectors", Pipe(VERTICAL_PIPE), []Direction{NORTH, SOUTH}},
		{"The horizontal pipe has east and west connectors", Pipe(HORIZONTAL_PIPE), []Direction{EAST, WEST}},
		{"The N-E bend has north and east connectors", Pipe(NORTH_TO_EAST_BEND), []Direction{NORTH, EAST}},
		{"The N-W bend has north and west connectors", Pipe(NORTH_TO_WEST_BEND), []Direction{NORTH, WEST}},
		{"The S-E bend has south and east connectors", Pipe(SOUTH_TO_EAST_BEND), []Direction{SOUTH, EAST}},
		{"The S-E bend has south and east connectors", Pipe(SOUTH_TO_WEST_BEND), []Direction{SOUTH, WEST}},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			answer := tt.pipe.ConnectionPoints()
			if !slices.Equal(answer, tt.expected) {
				t.Errorf("Got %v, expected %v (for pipe: %v)", answer, tt.expected, string(tt.pipe))
			}
		})
	}
}

package maze

import "testing"

func TestIsPipe(t *testing.T) {
	var testCases = []struct {
		testName string
		tile     Tile
		expected bool
	}{
		{"Ground is not a piple", GROUND, false},
		{"The entry point is a pipe", ENTRY, true},
		{"The N-S path is a pipe", VERTICAL_PIPE, true},
		{"The E-W path is a pipe", HORIZONTAL_PIPE, true},
		{"The N-E bend is a pipe", NORTH_TO_EAST_BEND, true},
		{"The N-W bend is a pipe", NORTH_TO_WEST_BEND, true},
		{"The S-E bend is a pipe", SOUTH_TO_EAST_BEND, true},
		{"The S-W bend is a pipe", SOUTH_TO_WEST_BEND, true},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			answer := tt.tile.IsPipe()
			if answer != tt.expected {
				t.Errorf("Got %v, expected %v (for tile: %v)", answer, tt.expected, tt.tile)
			}
		})
	}
}

package maze

import (
	"fmt"
	"testing"

	"amheklerior.com/advent-of-code-2023/utils"
)

func TestEntryPipeIdentification(t *testing.T) {
	terrain := BuildTerrain(utils.ReadFile("../data/entry-identification.test.txt"))
	fmt.Printf("terrain (%v x %v): %v", len(terrain), len(terrain[0]), terrain.String())

	testCases := []struct {
		testName      string
		entryPosition Position
		expected      Pipe
	}{
		{"Identify as SOUTH_TO_EAST_BEND", Position{1, 1}, Pipe(SOUTH_TO_EAST_BEND)},
		{"Identyfy as SOUTH_TO_WEST_BEND", Position{1, 9}, Pipe(SOUTH_TO_WEST_BEND)},
		{"Identify as VERTICAL_PIPE", Position{3, 1}, Pipe(VERTICAL_PIPE)},
		{"Identify as HORIZONTAL_PIPE", Position{2, 5}, Pipe(HORIZONTAL_PIPE)},
		{"Identify as NORTH_TO_EAST_BEND", Position{7, 1}, Pipe(NORTH_TO_EAST_BEND)},
		{"Identify as NORTH_TO_WEST_BEND", Position{7, 9}, Pipe(NORTH_TO_WEST_BEND)},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			pipe := IdentyfyEntryPipeType(terrain, tt.entryPosition)
			if pipe != tt.expected {
				t.Errorf("Got %v, expected %v! in position (%v-%v) ", pipe, tt.expected, tt.entryPosition.i, tt.entryPosition.j)
			}
		})
	}
}

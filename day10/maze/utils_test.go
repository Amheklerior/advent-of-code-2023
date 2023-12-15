package maze

import (
	"fmt"
	"testing"

	"amheklerior.com/advent-of-code-2023/utils"
)

func TestFindEntryPosition(t *testing.T) {
	testCases := []struct {
		testName string
		terrain  Terrain
		expected *Position
	}{
		{"Entry point is on top-right corner",
			Terrain{
				{GROUND, GROUND, ENTRY},
				{GROUND, GROUND, GROUND},
				{GROUND, GROUND, GROUND},
			},
			&Position{0, 2},
		},
		{"Entry point is on bottom-right corner",
			Terrain{
				{GROUND, GROUND, GROUND},
				{GROUND, GROUND, GROUND},
				{GROUND, GROUND, ENTRY},
			},
			&Position{2, 2},
		},
		{"Entry point is on top-left corner",
			Terrain{
				{ENTRY, GROUND, GROUND},
				{GROUND, GROUND, GROUND},
				{GROUND, GROUND, GROUND},
			},
			&Position{0, 0},
		},
		{"Entry point is on bottom-left corner",
			Terrain{
				{GROUND, GROUND, GROUND},
				{GROUND, GROUND, GROUND},
				{ENTRY, GROUND, GROUND},
			},
			&Position{2, 0},
		},
		{"Entry point is in the middle of the map",
			Terrain{
				{GROUND, GROUND, GROUND},
				{GROUND, ENTRY, GROUND},
				{GROUND, GROUND, GROUND},
			},
			&Position{1, 1},
		},
		{"Entry point not present",
			Terrain{
				{GROUND, GROUND, GROUND},
				{GROUND, GROUND, GROUND},
				{GROUND, GROUND, GROUND},
			},
			nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			answer := FindEntryPosition(&tt.terrain)
			if tt.expected == nil && answer != nil {
				t.Errorf("Got %v, expected to not be found", *answer)
			}
			if tt.expected != nil && *answer != *tt.expected {
				t.Errorf("Got %v, expected %v", *answer, *tt.expected)
			}
		})
	}
}

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
			pipe := IdentyfyEntryPipeType(&terrain, tt.entryPosition)
			if pipe != tt.expected {
				t.Errorf("Got %v, expected %v! in position (%v-%v) ", pipe, tt.expected, tt.entryPosition.i, tt.entryPosition.j)
			}
		})
	}
}

func TestFollowPipe(t *testing.T) {
	terrain := Terrain{
		{GROUND, GROUND, SOUTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND, GROUND},
		{GROUND, SOUTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, VERTICAL_PIPE, GROUND},
		{ENTRY, NORTH_TO_WEST_BEND, GROUND, NORTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND},
		{VERTICAL_PIPE, SOUTH_TO_EAST_BEND, HORIZONTAL_PIPE, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND},
		{NORTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, GROUND, GROUND, GROUND},
	}

	testCases := []struct {
		testName    string
		in          Position
		comingFrom  Position
		expectedPos Position
		expectedDir Direction
	}{
		{"Properly follows the pipe #1 (inside)", Position{2, 3}, Position{2, 4}, Position{1, 3}, NORTH},
		{"Properly follows the pipe #2 (inside)", Position{3, 2}, Position{3, 1}, Position{3, 3}, EAST},
		{"Properly follows the pipe #3 (inside)", Position{2, 1}, Position{1, 1}, Position{2, 0}, WEST},
		{"Properly follows the pipe #4 (on left boundary)", Position{4, 0}, Position{3, 0}, Position{4, 1}, EAST},
		{"Properly follows the pipe #5 (on right boundary)", Position{2, 4}, Position{2, 3}, Position{3, 4}, SOUTH},
		{"Properly follows the pipe #6 (on top boundary)", Position{0, 2}, Position{1, 2}, Position{0, 3}, EAST},
		{"Properly follows the pipe #6 (on bottom boundary)", Position{4, 0}, Position{4, 1}, Position{3, 0}, NORTH},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			pos, dir := FollowPipe(&terrain, tt.in, tt.comingFrom)
			if pos != tt.expectedPos || dir != tt.expectedDir {
				t.Errorf("Got %v/%v, expected %v/%v", pos, dir, tt.expectedPos, tt.expectedDir)
			}
		})
	}
}

func TestFollowPipeFromEntry(t *testing.T) {
	terrain := Terrain{
		{GROUND, HORIZONTAL_PIPE, ENTRY},
		{VERTICAL_PIPE, ENTRY, HORIZONTAL_PIPE},
		{GROUND, VERTICAL_PIPE, GROUND},
	}
	expected := Position{2, 1}

	answer, _ := FollowPipe(&terrain, Position{1, 1}, Position{1, 1})
	if answer != expected {
		t.Errorf("Got %v, expected %v", answer, expected)
	}
}

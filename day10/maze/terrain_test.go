package maze

import (
	"slices"
	"testing"

	"amheklerior.com/advent-of-code-2023/utils"
)

func TestTerrainBuilding(t *testing.T) {
	testCases := []struct {
		testName string
		file     string
		expected Terrain
	}{
		{
			"Build test #1",
			"../data/p1-input.test.txt",
			Terrain{
				{GROUND, GROUND, GROUND, GROUND, GROUND},
				{GROUND, ENTRY, HORIZONTAL_PIPE, SOUTH_TO_WEST_BEND, GROUND},
				{GROUND, VERTICAL_PIPE, GROUND, VERTICAL_PIPE, GROUND},
				{GROUND, NORTH_TO_EAST_BEND, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND, GROUND},
				{GROUND, GROUND, GROUND, GROUND, GROUND},
			},
		}, {
			"Build test #2",
			"../data/p2-input.test.txt",
			Terrain{
				{GROUND, GROUND, SOUTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND, GROUND},
				{GROUND, SOUTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, VERTICAL_PIPE, GROUND},
				{ENTRY, NORTH_TO_WEST_BEND, GROUND, NORTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND},
				{VERTICAL_PIPE, SOUTH_TO_EAST_BEND, HORIZONTAL_PIPE, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND},
				{NORTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, GROUND, GROUND, GROUND},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			input := utils.ReadFile(tt.file)
			answer := BuildTerrain(input)
			for i, row := range answer {
				if !slices.Equal(row, tt.expected[i]) {
					t.Errorf("Got %v, expected %v (for row: %v)", row, tt.expected[i], i)
				}
			}
		})
	}
}

func TestAdjacentPipeFinding(t *testing.T) {
	terrain := Terrain{
		{GROUND, GROUND, SOUTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND, GROUND},
		{GROUND, SOUTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, VERTICAL_PIPE, GROUND},
		{ENTRY, NORTH_TO_WEST_BEND, GROUND, NORTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND},
		{VERTICAL_PIPE, SOUTH_TO_EAST_BEND, HORIZONTAL_PIPE, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND},
		{NORTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, GROUND, GROUND, GROUND},
	}

	testCases := []struct {
		testName   string
		in         Position
		comingFrom Position
		expected   Position
	}{
		{"Properly follows the pipe #1 (inside)", Position{2, 3}, Position{2, 4}, Position{1, 3}},
		{"Properly follows the pipe #2 (inside)", Position{3, 2}, Position{3, 1}, Position{3, 3}},
		{"Properly follows the pipe #3 (inside)", Position{1, 1}, Position{1, 2}, Position{2, 1}},
		{"Properly follows the pipe #4 (on left boundary)", Position{4, 0}, Position{3, 0}, Position{4, 1}},
		{"Properly follows the pipe #5 (on right boundary)", Position{2, 4}, Position{2, 3}, Position{3, 4}},
		{"Properly follows the pipe #6 (on top boundary)", Position{0, 2}, Position{1, 2}, Position{0, 3}},
		{"Properly follows the pipe #6 (on bottom boundary)", Position{4, 0}, Position{4, 1}, Position{3, 0}},
		// {"Logs an error there's no pipe to follow in current position", Position{0, 0}, Position{0, 0}, Position{0, 0}},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			answer := terrain.FollowPipe(tt.in, tt.comingFrom)
			if answer != tt.expected {
				t.Errorf("Got %v, expected %v", answer, tt.expected)
			}
		})
	}
}

func TestPositionOfTile(t *testing.T) {

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
			answer := tt.terrain.PositionOfTile(ENTRY)
			if tt.expected == nil && answer != nil {
				t.Errorf("Got %v, expected to not be found", *answer)
			}
			if tt.expected != nil && *answer != *tt.expected {
				t.Errorf("Got %v, expected %v", *answer, *tt.expected)
			}
		})
	}
}

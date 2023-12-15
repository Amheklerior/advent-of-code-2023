package maze

import (
	"slices"
	"testing"
)

func TestBuildLoopPath(t *testing.T) {
	testCases := []struct {
		testName string
		terrain  Terrain
		expected LoopPath
	}{
		{
			"Build test #1 (simple ring loop)",
			Terrain{
				{GROUND, GROUND, GROUND, GROUND, GROUND},
				{GROUND, ENTRY, HORIZONTAL_PIPE, SOUTH_TO_WEST_BEND, GROUND},
				{GROUND, VERTICAL_PIPE, GROUND, VERTICAL_PIPE, GROUND},
				{GROUND, NORTH_TO_EAST_BEND, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND, GROUND},
				{GROUND, GROUND, GROUND, GROUND, GROUND},
			},
			LoopPath{
				Position{1, 1},
				Position{2, 1},
				Position{3, 1},
				Position{3, 2},
				Position{3, 3},
				Position{2, 3},
				Position{1, 3},
				Position{1, 2},
				Position{1, 1},
			},
		}, {
			"Build test #2 (longer convoluted loop)",
			Terrain{
				{GROUND, GROUND, SOUTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND, GROUND},
				{GROUND, SOUTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, VERTICAL_PIPE, GROUND},
				{ENTRY, NORTH_TO_WEST_BEND, GROUND, NORTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND},
				{VERTICAL_PIPE, SOUTH_TO_EAST_BEND, HORIZONTAL_PIPE, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND},
				{NORTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, GROUND, GROUND, GROUND},
			},
			LoopPath{
				Position{2, 0},
				Position{3, 0},
				Position{4, 0},
				Position{4, 1},
				Position{3, 1},
				Position{3, 2},
				Position{3, 3},
				Position{3, 4},
				Position{2, 4},
				Position{2, 3},
				Position{1, 3},
				Position{0, 3},
				Position{0, 2},
				Position{1, 2},
				Position{1, 1},
				Position{2, 1},
				Position{2, 0},
			},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			answer := BuildPipeLoop(&tt.terrain)
			if !slices.Equal(answer, tt.expected) {
				t.Errorf("Got %v, expected %v", answer, tt.expected)
			}
		})
	}
}

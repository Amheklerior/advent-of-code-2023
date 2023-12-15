package maze

import (
	"slices"
	"testing"

	"amheklerior.com/advent-of-code-2023/utils"
)

func TestBuildTerrain(t *testing.T) {
	testCases := []struct {
		testName string
		file     string
		expected Terrain
	}{
		{
			"Build test #1",
			"../data/shortpipe.test.txt",
			Terrain{
				{GROUND, GROUND, GROUND, GROUND, GROUND},
				{GROUND, ENTRY, HORIZONTAL_PIPE, SOUTH_TO_WEST_BEND, GROUND},
				{GROUND, VERTICAL_PIPE, GROUND, VERTICAL_PIPE, GROUND},
				{GROUND, NORTH_TO_EAST_BEND, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND, GROUND},
				{GROUND, GROUND, GROUND, GROUND, GROUND},
			},
		}, {
			"Build test #2",
			"../data/longpipe.test.txt",
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
			if !answer.Equal(tt.expected) {
				t.Errorf(
					"Got %v, expected %v",
					answer.String(),
					tt.expected.String(),
				)
			}
		})
	}
}

func TestForEach(t *testing.T) {
	terrain := Terrain{
		{GROUND, GROUND, GROUND},
		{GROUND, GROUND, GROUND},
		{GROUND, GROUND, GROUND},
	}
	expected := Terrain{
		{VERTICAL_PIPE, VERTICAL_PIPE, VERTICAL_PIPE},
		{VERTICAL_PIPE, VERTICAL_PIPE, VERTICAL_PIPE},
		{VERTICAL_PIPE, VERTICAL_PIPE, VERTICAL_PIPE},
	}

	answer := terrain
	answer.ForEach(func(tile Tile, i, j int) {
		answer[i][j] = VERTICAL_PIPE
	})

	if !answer.Equal(expected) {
		t.Errorf(
			"Got %v, expected %v",
			answer.String(),
			expected.String(),
		)
	}
}

func TestEqual(t *testing.T) {
	t1 := Terrain{
		{GROUND, GROUND, GROUND},
		{GROUND, GROUND, GROUND},
		{GROUND, GROUND, GROUND},
	}
	t2 := t1
	t3 := Terrain{
		{VERTICAL_PIPE, VERTICAL_PIPE, VERTICAL_PIPE},
		{VERTICAL_PIPE, VERTICAL_PIPE, VERTICAL_PIPE},
		{VERTICAL_PIPE, VERTICAL_PIPE, VERTICAL_PIPE},
	}

	if !t1.Equal(t2) {
		t.Errorf(
			"Got false, expected to be equal. \n%v\n%v",
			t1.String(),
			t2.String(),
		)
	}
	if t1.Equal(t3) {
		t.Errorf(
			"Got true, expected to not be equal. \n%v\n%v",
			t1.String(),
			t3.String(),
		)
	}
}

func TestCleanupFromJunkPipes(t *testing.T) {
	terrain := Terrain{
		{VERTICAL_PIPE, HORIZONTAL_PIPE, SOUTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND, SOUTH_TO_EAST_BEND},
		{GROUND, SOUTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, VERTICAL_PIPE, VERTICAL_PIPE},
		{ENTRY, NORTH_TO_WEST_BEND, HORIZONTAL_PIPE, NORTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND},
		{VERTICAL_PIPE, SOUTH_TO_EAST_BEND, HORIZONTAL_PIPE, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND},
		{NORTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, GROUND, NORTH_TO_EAST_BEND, NORTH_TO_WEST_BEND},
	}
	expected := Terrain{
		{GROUND, GROUND, SOUTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND, GROUND},
		{GROUND, SOUTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, VERTICAL_PIPE, GROUND},
		{ENTRY, NORTH_TO_WEST_BEND, GROUND, NORTH_TO_EAST_BEND, SOUTH_TO_WEST_BEND},
		{VERTICAL_PIPE, SOUTH_TO_EAST_BEND, HORIZONTAL_PIPE, HORIZONTAL_PIPE, NORTH_TO_WEST_BEND},
		{NORTH_TO_EAST_BEND, NORTH_TO_WEST_BEND, GROUND, GROUND, GROUND},
	}
	loop := BuildPipeLoop(&terrain)

	terrain.CleanupFromJunkPipes(loop)
	for i := range terrain {
		if !slices.Equal(terrain[i], expected[i]) {
			t.Errorf("Got %v, expected %v (at row %v)", terrain[i], expected[i], i)
		}
	}
}

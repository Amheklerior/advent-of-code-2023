package maze

import (
	"testing"

	"amheklerior.com/advent-of-code-2023/utils"
)

func TestFSMSolver(t *testing.T) {
	testCases := []struct {
		testName string
		filePath string
		expected int
	}{
		{"Only enclosed tiles are counted (simple)", "../data/enclosed-simple.test.txt", 4},
		{"Tiles that visually seem to be encolsed are properly ignored", "../data/enclosed-tricky.test.txt", 4},
		{"Only enclosed tiles are counted (complex 1)", "../data/enclosed-complex.test.txt", 8},
		{"Only enclosed tiles are counted (complex 2)", "../data/p2-input.test.txt", 10},
	}

	for _, tt := range testCases {
		t.Run(tt.testName, func(t *testing.T) {
			input := utils.ReadFile(tt.filePath)
			terrain := BuildTerrain(input)
			loop := BuildPipeLoop(&terrain)
			terrain.CleanupFromJunkPipes(loop)
			fsm := NewFSM(terrain)
			answer := fsm.Solve()
			if answer != tt.expected {
				t.Errorf("Got %v, expected %v (for terrain in '%v')", answer, tt.expected, tt.filePath)
			}
		})
	}
}

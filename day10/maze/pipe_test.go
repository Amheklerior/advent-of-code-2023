package maze

import (
	"fmt"
	"math/rand"
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

func TestConnectWith(t *testing.T) {
	// var testCases = []struct {
	// 	testName string
	// 	this, other     Pipe
	// 	dir Direction
	// 	expected bool
	// }{
	// 	{"The vertical pipe has north and south connectors", Pipe(VERTICAL_PIPE), , true},
	// }

	// s := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(42)         // create a random generator with a given seed
	// r := rand.New(source)     // random generators with the same seed produce the same sequence of random numbers
	// r := rand.New(s)          // if you want a non deterministic sequence, feed a source with a changing seed

	// r.Intn(100)               // rand.Rand exposes the same func as the rand package
	// r.Perm(3)                 // a random permutation of the numbers between 0 and i • ex. [1, 2, 0]

	pipes := []Pipe{
		Pipe(HORIZONTAL_PIPE),
		Pipe(VERTICAL_PIPE),
		Pipe(SOUTH_TO_EAST_BEND),
		Pipe(SOUTH_TO_WEST_BEND),
		Pipe(NORTH_TO_EAST_BEND),
		Pipe(NORTH_TO_WEST_BEND),
	}

	directions := []Direction{
		NORTH,
		SOUTH,
		EAST,
		WEST,
	}

	for i := 0; i < 50; i++ {
		p := rand.Intn(len(pipes))
		o := rand.Intn(len(pipes))
		d := rand.Intn(len(directions))
		pipe := pipes[p]
		other := pipes[o]
		dir := directions[d]
		t.Run(
			fmt.Sprintf(
				"Test #%v: from pipe %v to pipe %v along dir %v",
				i,
				string(pipe),
				string(other),
				string(dir),
			), func(t *testing.T) {
				answer := pipe.CanConnectWith(other, dir)
				fmt.Printf("from %v to %v along %v --> %v", string(pipe), string(other), string(dir), answer)
				// if answer !=  {
				// 	t.Errorf("Got %v, expected %v (for pipe: %v)", answer, tt.expected, string(tt.pipe))
				// }
			})
	}
}

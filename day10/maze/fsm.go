package maze

import (
	"fmt"
)

type State rune

const (
	OUTSIDE  State = 'O'
	INSIDE   State = 'I'
	ENTERING State = 'E'
	EXITING  State = 'X'
)

type FSM struct {
	currentState      State
	Count             int
	terrain           Terrain
	transitioningWith *Pipe
}

func NewFSM(terrain Terrain) FSM {
	var fsm = FSM{
		currentState:      OUTSIDE,
		Count:             0,
		terrain:           terrain,
		transitioningWith: new(Pipe),
	}
	return fsm
}

func (fsm *FSM) transitTo(state State, pipe *Pipe) {
	fsm.currentState = state
	if pipe != nil {
		fsm.transitioningWith = pipe
	}
}

func (fsm *FSM) Solve() int {
	fsm.terrain.ForEach(func(tile Tile, i, j int) {
		fsm.Process(Position{i, j})
	})
	return fsm.Count
}

func (fsm *FSM) Process(pos Position) {
	tile := fsm.terrain.At(pos)

	switch fsm.currentState {
	case OUTSIDE:
		if tile == GROUND {
			// Ignore all ground and junk pipes
			return
		}
		if tile == ENTRY {
			tile = Tile(IdentyfyEntryPipeType(&fsm.terrain, pos))
		}
		if tile == VERTICAL_PIPE {
			// vertical pipes always count as crossing
			fsm.transitTo(INSIDE, nil)
			return
		}
		if tile == SOUTH_TO_EAST_BEND || tile == NORTH_TO_EAST_BEND {
			// Crossing that must be confirmed by a later piece
			pipe := Pipe(tile)
			fsm.transitTo(ENTERING, &pipe)
			return
		}
		// walking from left to right, we can never encounter
		// NORTH_TO_WEST or SOUTH_TO_WEST pipes coming from outside
	case INSIDE:
		// fmt.Printf("Processing INSIDE state...\n")
		if tile == GROUND {
			// Count all ground and junk pipes
			fsm.Count++
			return
		}
		if tile == ENTRY {
			tile = Tile(IdentyfyEntryPipeType(&fsm.terrain, pos))
		}
		if tile == VERTICAL_PIPE {
			// vertical pipes always count as crossing
			fsm.transitTo(OUTSIDE, nil)
			return
		}
		if tile == SOUTH_TO_EAST_BEND || tile == NORTH_TO_EAST_BEND {
			// crossing that must be confirmed by a later piece
			pipe := Pipe(tile)
			fsm.transitTo(EXITING, &pipe)
			return
		}
	case ENTERING:
		if tile == ENTRY {
			tile = Tile(IdentyfyEntryPipeType(&fsm.terrain, pos))
		}
		if (*fsm.transitioningWith == Pipe(NORTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) ||
			(*fsm.transitioningWith == Pipe(SOUTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) {
			fsm.transitTo(INSIDE, new(Pipe))
			// crossing confirmed
			return
		}
		if (*fsm.transitioningWith == Pipe(NORTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) ||
			(*fsm.transitioningWith == Pipe(SOUTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) {
			fsm.transitTo(OUTSIDE, new(Pipe))
			// crossing discarded
			return
		}
	case EXITING:
		if tile == ENTRY {
			tile = Tile(IdentyfyEntryPipeType(&fsm.terrain, pos))
		}
		if (*fsm.transitioningWith == Pipe(NORTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) ||
			(*fsm.transitioningWith == Pipe(SOUTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) {
			fsm.transitTo(OUTSIDE, new(Pipe))
			// crossing confirmed
			return
		}
		if (*fsm.transitioningWith == Pipe(NORTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) ||
			(*fsm.transitioningWith == Pipe(SOUTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) {
			fsm.transitTo(INSIDE, new(Pipe))
			// crossing discarded
			return
		}
	}
}

func (fsm *FSM) String() string {
	return fmt.Sprintf(
		"fsm: {\n\tcurrent: %v \n\tcount: %v \n\ttransitioningWith: %v\n}\n",
		string(fsm.currentState),
		fsm.Count,
		string(*fsm.transitioningWith),
	)
}

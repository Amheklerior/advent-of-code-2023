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
	loop              LoopPath
	transitioningWith *Pipe
}

func NewFSM(terrain Terrain, loop LoopPath) FSM {
	return FSM{
		currentState:      OUTSIDE,
		Count:             0,
		terrain:           terrain,
		loop:              loop,
		transitioningWith: new(Pipe),
	}
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

	// fmt.Printf("%vprocessing tile %v in position (%v,%v)\n", fsm, string(tile), position.i, position.j)
	switch fsm.currentState {
	case OUTSIDE:
		// fmt.Printf("Processing OUTSIDE state...\n")
		if !fsm.loop.Contains(pos) {
			// fmt.Printf("%v is not part of the loop\n", string(tile))
			return
		}
		if tile == ENTRY {
			tile = Tile(IdentyfyEntryPipeType(&fsm.terrain, pos))
		}
		if tile == VERTICAL_PIPE {
			fsm.transitTo(INSIDE, nil)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
		if tile == SOUTH_TO_EAST_BEND || tile == NORTH_TO_EAST_BEND {
			pipe := Pipe(tile)
			fsm.transitTo(ENTERING, &pipe)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
	case INSIDE:
		// fmt.Printf("Processing INSIDE state...\n")
		if tile == GROUND || !fsm.loop.Contains(pos) {
			fsm.Count++
			// fmt.Printf("Updated count to %v in position (%v,%v)\n", fsm.Count, position.i, position.j)
			return
		}
		if tile == ENTRY {
			tile = Tile(IdentyfyEntryPipeType(&fsm.terrain, pos))
		}
		if tile == VERTICAL_PIPE {
			fsm.transitTo(OUTSIDE, nil)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
		if tile == SOUTH_TO_EAST_BEND || tile == NORTH_TO_EAST_BEND {
			pipe := Pipe(tile)
			fsm.transitTo(EXITING, &pipe)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
	case ENTERING:
		// fmt.Printf("Processing ENTERING state...\n")
		if (*fsm.transitioningWith == Pipe(NORTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) ||
			(*fsm.transitioningWith == Pipe(SOUTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) {
			fsm.transitTo(INSIDE, new(Pipe))
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
		if (*fsm.transitioningWith == Pipe(NORTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) ||
			(*fsm.transitioningWith == Pipe(SOUTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) {
			fsm.transitTo(OUTSIDE, new(Pipe))
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
	case EXITING:
		// fmt.Printf("Processing EXITING state...\n\n")
		if (*fsm.transitioningWith == Pipe(NORTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) ||
			(*fsm.transitioningWith == Pipe(SOUTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) {
			fsm.transitTo(OUTSIDE, new(Pipe))
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
		if (*fsm.transitioningWith == Pipe(NORTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) ||
			(*fsm.transitioningWith == Pipe(SOUTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) {
			fsm.transitTo(INSIDE, new(Pipe))
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
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

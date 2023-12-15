package maze

import "fmt"

type State rune

const (
	OUTSIDE  State = 'O'
	INSIDE   State = 'I'
	ENTERING State = 'E'
	EXITING  State = 'X'
)

type FSM struct {
	currentState State
	Count        int
	terrain      Terrain
	loop         PipeLoop
	enteringWith *Pipe
	exitingWith  *Pipe
}

func NewFSM(terrain Terrain, loop PipeLoop) FSM {
	return FSM{
		currentState: OUTSIDE,
		Count:        0,
		terrain:      terrain,
		loop:         loop,
		enteringWith: new(Pipe),
		exitingWith:  new(Pipe),
	}
}

func (fsm *FSM) transitTo(state State) {
	fsm.currentState = state
}

func (fsm *FSM) IsPartOfTheLoop(tile Tile, pos Position) bool {
	return fsm.loop.Contains(LoopPortion{Pipe(tile), pos})
}

func (fsm *FSM) Process(position Position) {
	tile := fsm.terrain.At(position)

	// fmt.Printf("%vprocessing tile %v in position (%v,%v)\n", fsm, string(tile), position.i, position.j)
	switch fsm.currentState {
	case OUTSIDE:
		// fmt.Printf("Processing OUTSIDE state...\n")
		if !fsm.IsPartOfTheLoop(tile, position) {
			// fmt.Printf("%v is not part of the loop\n", string(tile))
			return
		}
		if tile == ENTRY {
			tile = Tile(IdentyfyEntryPipeType(fsm.terrain, position))
		}
		if tile == VERTICAL_PIPE {
			fsm.transitTo(INSIDE)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
		if tile == SOUTH_TO_EAST_BEND || tile == NORTH_TO_EAST_BEND {
			*fsm.enteringWith = Pipe(tile)
			fsm.transitTo(ENTERING)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
	case INSIDE:
		// fmt.Printf("Processing INSIDE state...\n")
		if tile == GROUND || !fsm.IsPartOfTheLoop(tile, position) {
			fsm.Count++
			// fmt.Printf("Updated count to %v in position (%v,%v)\n", fsm.Count, position.i, position.j)
			return
		}
		if tile == ENTRY {
			tile = Tile(IdentyfyEntryPipeType(fsm.terrain, position))
		}
		if tile == VERTICAL_PIPE {
			fsm.transitTo(OUTSIDE)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
		if tile == SOUTH_TO_EAST_BEND || tile == NORTH_TO_EAST_BEND {
			*fsm.exitingWith = Pipe(tile)
			fsm.transitTo(EXITING)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
	case ENTERING:
		// fmt.Printf("Processing ENTERING state...\n")
		if (*fsm.enteringWith == Pipe(NORTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) ||
			(*fsm.enteringWith == Pipe(SOUTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) {
			fsm.enteringWith = new(Pipe)
			fsm.transitTo(INSIDE)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
		if (*fsm.enteringWith == Pipe(NORTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) ||
			(*fsm.enteringWith == Pipe(SOUTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) {
			fsm.enteringWith = new(Pipe)
			fsm.transitTo(OUTSIDE)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
	case EXITING:
		// fmt.Printf("Processing EXITING state...\n\n")
		if (*fsm.exitingWith == Pipe(NORTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) ||
			(*fsm.exitingWith == Pipe(SOUTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) {
			fsm.exitingWith = new(Pipe)
			fsm.transitTo(OUTSIDE)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
		if (*fsm.exitingWith == Pipe(NORTH_TO_EAST_BEND) && tile == NORTH_TO_WEST_BEND) ||
			(*fsm.exitingWith == Pipe(SOUTH_TO_EAST_BEND) && tile == SOUTH_TO_WEST_BEND) {
			fsm.exitingWith = new(Pipe)
			fsm.transitTo(INSIDE)
			// fmt.Printf("Change fsm state to %v \n", fsm.currentState)
			return
		}
	}
}

func (fsm *FSM) Solve() int {
	for i := range fsm.terrain {
		for j := range fsm.terrain[i] {
			fsm.Process(Position{i, j})
		}
	}
	return fsm.Count
}

func (fsm *FSM) String() string {
	return fmt.Sprintf("fsm: {\n\tcurrent: %v \n\tcount: %v \n\tenteringWith: %v \n\texitingWith: %v\n}\n", string(fsm.currentState), fsm.Count, string(*fsm.enteringWith), string(*fsm.exitingWith))
}

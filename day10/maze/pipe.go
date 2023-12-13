package maze

import (
	"log"
	"slices"
)

type Pipe Tile

var connectionsFor map[Pipe][]Direction = map[Pipe][]Direction{
	Pipe(ENTRY):              {WEST, EAST, SOUTH, NORTH},
	Pipe(VERTICAL_PIPE):      {NORTH, SOUTH},
	Pipe(HORIZONTAL_PIPE):    {EAST, WEST},
	Pipe(NORTH_TO_EAST_BEND): {NORTH, EAST},
	Pipe(NORTH_TO_WEST_BEND): {NORTH, WEST},
	Pipe(SOUTH_TO_EAST_BEND): {SOUTH, EAST},
	Pipe(SOUTH_TO_WEST_BEND): {SOUTH, WEST},
}

func (pipe Pipe) ConnectionPoints() []Direction {
	return connectionsFor[pipe]
}

func (pipe Pipe) HasConnectorAlong(dir Direction) bool {
	return slices.Contains(pipe.ConnectionPoints(), dir)
}

func (pipe Pipe) CanConnectWith(otherPipe Pipe, dir Direction) bool {
	if !pipe.HasConnectorAlong(dir) {
		log.Fatalf("Pipe %v cannot connect along %v. %v is not one of its connection points.", pipe, dir, dir)
	}
	for _, otherPipeConnector := range otherPipe.ConnectionPoints() {
		if dir.connectsTo(otherPipeConnector) {
			return true
		}
	}
	return false
}

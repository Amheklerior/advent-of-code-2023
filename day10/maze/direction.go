package maze

type Direction rune
type Vector struct {
	i, j int
}

const (
	NORTH Direction = 'N'
	SOUTH Direction = 'S'
	WEST  Direction = 'W'
	EAST  Direction = 'E'
)

var vector map[Direction]Vector = map[Direction]Vector{
	NORTH: {-1, 0},
	SOUTH: {1, 0},
	EAST:  {0, 1},
	WEST:  {0, -1},
}

var adjacentOf map[Direction]Direction = map[Direction]Direction{
	NORTH: SOUTH,
	SOUTH: NORTH,
	EAST:  WEST,
	WEST:  EAST,
}

func (dir Direction) connectsTo(other Direction) bool {
	return adjacentOf[dir] == other
}

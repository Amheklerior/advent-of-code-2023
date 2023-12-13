package maze

type Tile rune

const (
	GROUND             Tile = '.'
	ENTRY              Tile = 'S'
	VERTICAL_PIPE      Tile = '|'
	HORIZONTAL_PIPE    Tile = '-'
	NORTH_TO_EAST_BEND Tile = 'L'
	NORTH_TO_WEST_BEND Tile = 'J'
	SOUTH_TO_EAST_BEND Tile = 'F'
	SOUTH_TO_WEST_BEND Tile = '7'
)

func (t Tile) IsPipe() bool {
	return t != GROUND
}

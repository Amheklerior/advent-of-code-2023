package maze

import (
	"log"
)

func FindEntryPosition(t *Terrain) *Position {
	for i, row := range *t {
		for j, tile := range row {
			if tile == ENTRY {
				return &Position{i, j}
			}
		}
	}
	return nil
}

func IdentyfyEntryPipeType(t *Terrain, pos Position) Pipe {
	if t.At(pos) != ENTRY {
		log.Fatalf("tile at position %v is not the ENTRY tile. Found %v instead.", pos, string(t.At(pos)))
	}

	firstConnection, firstDir := FollowPipe(t, pos, pos)
	_, secondDir := FollowPipe(t, pos, firstConnection)

	if firstDir == NORTH && secondDir == SOUTH || firstDir == SOUTH && secondDir == NORTH {
		return Pipe(VERTICAL_PIPE)
	}
	if firstDir == EAST && secondDir == WEST || firstDir == WEST && secondDir == EAST {
		return Pipe(HORIZONTAL_PIPE)
	}
	if firstDir == NORTH && secondDir == EAST || firstDir == EAST && secondDir == NORTH {
		return Pipe(NORTH_TO_EAST_BEND)
	}
	if firstDir == NORTH && secondDir == WEST || firstDir == WEST && secondDir == NORTH {
		return Pipe(NORTH_TO_WEST_BEND)
	}
	if firstDir == SOUTH && secondDir == EAST || firstDir == EAST && secondDir == SOUTH {
		return Pipe(SOUTH_TO_EAST_BEND)
	}

	return Pipe(SOUTH_TO_WEST_BEND)
}

func FollowPipe(t *Terrain, currentPipePosition, comingFrom Position) (Position, Direction) {
	if !t.At(currentPipePosition).IsPipe() {
		log.Fatalf("Found %v at position %v, which is not a pipe.", t.At(currentPipePosition), currentPipePosition)
	}

	onTopRow := currentPipePosition.i == 0
	onBottomRow := currentPipePosition.i == t.Height()-1
	onFirstCol := currentPipePosition.j == 0
	onLastCol := currentPipePosition.j == t.Width()-1

	var nextPipePos Position
	var onDir Direction
	currentPipe := Pipe(t.At(currentPipePosition))

	for _, dir := range currentPipe.ConnectionPoints() {
		if (dir == NORTH && onTopRow) ||
			(dir == SOUTH && onBottomRow) ||
			(dir == EAST && onLastCol) ||
			(dir == WEST && onFirstCol) {
			continue
		}

		positionToCheck := Position{
			currentPipePosition.i + vector[dir].i,
			currentPipePosition.j + vector[dir].j,
		}
		tile := t.At(positionToCheck)

		if comingFrom == positionToCheck || !tile.IsPipe() {
			continue
		}

		if currentPipe.CanConnectWith(Pipe(tile), dir) {
			nextPipePos = positionToCheck
			onDir = dir
		}
	}
	return nextPipePos, onDir
}

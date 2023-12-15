package maze

func IdentyfyEntryPipeType(terrain Terrain, pos Position) Pipe {
	firstConnection, firstDir := terrain.FollowPipe(pos, pos)
	_, secondDir := terrain.FollowPipe(pos, firstConnection)

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

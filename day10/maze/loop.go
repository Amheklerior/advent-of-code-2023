package maze

import (
	"slices"
)

type LoopPath []Position

func BuildPipeLoop(t *Terrain) LoopPath {
	var loop LoopPath
	entrypoint := FindEntryPosition(t)
	currPos, prevPos := *entrypoint, *entrypoint
	loop.Add(currPos)

	for t.At(currPos) != ENTRY || len(loop) <= 1 {
		nextPos, _ := FollowPipe(t, currPos, prevPos)
		loop.Add(nextPos)
		prevPos, currPos = currPos, nextPos
	}

	return loop
}

func (loop *LoopPath) Add(pos Position) {
	*loop = append(*loop, pos)
}

func (loop *LoopPath) Contains(pos Position) bool {
	return slices.Contains(*loop, pos)
}

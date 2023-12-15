package maze

import "slices"

type LoopPortion struct {
	pipe Pipe
	pos  Position
}

type PipeLoop []LoopPortion

func (loop *PipeLoop) Add(pipe Pipe, position Position) {
	*loop = append([]LoopPortion(*loop), LoopPortion{pipe, position})
}

func (loop *PipeLoop) Contains(portion LoopPortion) bool {
	return slices.Contains(*loop, portion)
}

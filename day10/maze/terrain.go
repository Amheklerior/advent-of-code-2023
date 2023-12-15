package maze

import (
	"fmt"
	"log"
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

type Terrain [][]Tile
type Position struct {
	i, j int
}

func BuildTerrain(input string) Terrain {
	var terrain Terrain
	scanner := utils.Scanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		terrain = append(terrain, []Tile(line))
	}
	return terrain
}

func (t *Terrain) At(p Position) Tile {
	return (*t)[p.i][p.j]
}

func (t *Terrain) Height() int {
	return len(*t)
}

func (t *Terrain) Width() int {
	if t.Height() == 0 {
		return 0
	}
	return len((*t)[0])
}

func (t *Terrain) PositionOfTile(tile Tile) *Position {
	for i, row := range [][]Tile(*t) {
		for j, item := range row {
			if item == tile {
				return &Position{i, j}
			}
		}
	}
	return nil
}

func (terrain *Terrain) FollowPipe(currentPipePosition, comingFrom Position) (Position, Direction) {
	if !terrain.At(currentPipePosition).IsPipe() {
		log.Fatalf("Found %v at position %v, which is not a pipe.", terrain.At(currentPipePosition), currentPipePosition)
	}

	onTopRow := currentPipePosition.i == 0
	onBottomRow := currentPipePosition.i == terrain.Height()-1
	onFirstCol := currentPipePosition.j == 0
	onLastCol := currentPipePosition.j == terrain.Width()-1

	var nextPipePos Position
	var onDir Direction
	currentPipe := Pipe(terrain.At(currentPipePosition))
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
		tile := terrain.At(positionToCheck)

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

func (t *Terrain) BuildPipeLoop() PipeLoop {
	var loop PipeLoop
	entrypoint := t.PositionOfTile(ENTRY)
	currPos, prevPos := *entrypoint, *entrypoint
	loop.Add(Pipe(ENTRY), currPos)

	for t.At(currPos) != ENTRY || len(loop) <= 1 {
		nextPos, _ := t.FollowPipe(currPos, prevPos)
		pipe := Pipe(t.At(nextPos))
		loop.Add(pipe, nextPos)
		prevPos, currPos = currPos, nextPos
	}

	return loop
}

func (t *Terrain) Cleanup(loop PipeLoop) {
	for i := range *t {
		for j := range (*t)[i] {
			pos := Position{i, j}
			tile := t.At(pos)
			if !tile.IsPipe() || !loop.Contains(LoopPortion{Pipe(tile), pos}) {
				(*t)[i][j] = GROUND
			}
		}
	}
}

func (t *Terrain) String() string {
	var builder []string
	builder = append(builder, fmt.Sprintf("T (%vx%v): [\n", (*t).Height(), (*t).Width()))
	for _, row := range *t {
		builder = append(builder, "\t[")
		for j, item := range row {
			builder = append(builder, string(item))
			if j != (*t).Width()-1 {
				builder = append(builder, " ")
			}
		}
		builder = append(builder, "]\n")
	}
	builder = append(builder, "]")

	return strings.Join(builder, "")
}

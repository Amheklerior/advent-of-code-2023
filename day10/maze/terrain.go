package maze

import (
	"fmt"
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

func (t *Terrain) CleanupFromJunkPipes(loop LoopPath) {
	t.ForEach(func(tile Tile, i int, j int) {
		if !tile.IsPipe() || !loop.Contains(Position{i, j}) {
			(*t)[i][j] = GROUND
		}
	})
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

func (t *Terrain) ForEach(callback func(Tile, int, int)) {
	for i := range *t {
		for j := range (*t)[i] {
			callback((*t)[i][j], i, j)
		}
	}
}

func (t *Terrain) Equal(other Terrain) bool {
	isEqual := true
	t.ForEach(func(tile Tile, i, j int) {
		if tile != other[i][j] {
			isEqual = false
		}
	})
	return isEqual
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

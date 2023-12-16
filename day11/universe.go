package day11

import (
	"math"
	"slices"

	"amheklerior.com/advent-of-code-2023/utils"
)

type Space utils.Matrix[rune]

type Coordinate struct {
	x, y int
}

type Universe struct {
	Space     Space
	emptyRows []int
	emptyCols []int
}

const (
	EMPTY  rune = '.'
	GALAXY rune = '#'
)

func NewUniverse(input string) Universe {
	var space Space
	scanner := utils.Scanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		space = append(space, []rune(line))
	}

	rows, cols := getEmptyAreas(space)
	return Universe{
		Space:     space,
		emptyRows: rows,
		emptyCols: cols,
	}
}

func (univ *Universe) GetGalaxies() []Coordinate {
	var galaxies []Coordinate

	u := utils.Matrix[rune](univ.Space)
	u.ForEach(func(space rune, i, j int) {
		if space == GALAXY {
			galaxies = append(galaxies, Coordinate{j, i})
		}
	})

	return galaxies
}

func getEmptyAreas(space Space) ([]int, []int) {
	var emptyRows, emptyCols []int
	var rowsWithGalaxies, colsWithGalaxies []int

	s := utils.Matrix[rune](space)
	s.ForEach(func(r rune, i, j int) {
		if r == GALAXY {
			rowsWithGalaxies = append(rowsWithGalaxies, i)
			colsWithGalaxies = append(colsWithGalaxies, j)
		}
	})

	s.ForEach(func(r rune, i, j int) {
		if !slices.Contains(rowsWithGalaxies, i) && !slices.Contains(emptyRows, i) {
			emptyRows = append(emptyRows, i)
		}
		if !slices.Contains(colsWithGalaxies, j) && !slices.Contains(emptyCols, j) {
			emptyCols = append(emptyCols, j)
		}
	})

	return emptyRows, emptyCols
}

func (univ *Universe) DistanceBetween(from, to Coordinate, expansionFactor int) int {
	verticalDistance := int(math.Abs(float64(to.y - from.y)))
	horizontalDistance := int(math.Abs(float64(to.x - from.x)))
	if expansionFactor > 1 {
		expansionFactor--
	}

	for _, idx := range univ.emptyCols {
		if (idx < from.x && idx > to.x) || (idx > from.x && idx < to.x) {
			verticalDistance += expansionFactor
		}
	}

	for _, idx := range univ.emptyRows {
		if (idx < from.y && idx > to.y) || (idx > from.y && idx < to.y) {
			verticalDistance += expansionFactor
		}
	}

	return verticalDistance + horizontalDistance
}

func (u *Universe) Solve(expFactor int) int {
	galaxies := u.GetGalaxies()
	sum := 0
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += u.DistanceBetween(galaxies[i], galaxies[j], expFactor)
		}
	}
	return sum
}

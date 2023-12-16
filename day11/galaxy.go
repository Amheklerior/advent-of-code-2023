package day11

import (
	"math"

	"amheklerior.com/advent-of-code-2023/utils"
)

type Space [][]rune

const (
	EMPTY  rune = '.'
	GALAXY rune = '#'
)

func NewSpace(input string) Space {
	var space Space
	scanner := utils.Scanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		space = append(space, []rune(line))
	}
	return space
}

func (s *Space) At(i, j int) rune {
	return (*s)[i][j]
}

func (s *Space) Height() int {
	return len(*s)
}

func (s *Space) Width() int {
	if s.Height() == 0 {
		return 0
	}
	return len((*s)[0])
}

func (s *Space) ForEach(callback func(rune, int, int)) {
	for i := range *s {
		for j := range (*s)[i] {
			callback((*s)[i][j], i, j)
		}
	}
}

func EmptyRow(length int) []rune {
	row := make([]rune, 0, length)
	for i := 0; i < length; i++ {
		row = append(row, EMPTY)
	}
	return row
}

func (s *Space) Expand() Space {
	colsWithGalaxies := make([]bool, s.Width(), s.Width())
	rowsWithGalaxies := make([]bool, s.Height(), s.Height())

	s.ForEach(func(s rune, i, j int) {
		if s == GALAXY {
			rowsWithGalaxies[i] = true
			colsWithGalaxies[j] = true
		}
	})

	addedCols := 0
	for i := range colsWithGalaxies {
		if !colsWithGalaxies[i] {
			addedCols++
		}
	}

	var expandedSpace Space
	addedRows := 0
	for i := range *s {
		if !rowsWithGalaxies[i] {
			expandedSpace = append(expandedSpace, EmptyRow(s.Width()+addedCols))
			addedRows++
		}

		expandedSpace = append(expandedSpace, make([]rune, 0))
		for j := range (*s)[i] {
			if !colsWithGalaxies[j] {
				expandedSpace[i+addedRows] = append(expandedSpace[i+addedRows], EMPTY)
			}
			expandedSpace[i+addedRows] = append(expandedSpace[i+addedRows], s.At(i, j))
		}
	}

	return expandedSpace
}

type Coordinate struct {
	x, y int
}

func (s *Space) GalaxiesMap() []Coordinate {
	var galaxies []Coordinate
	s.ForEach(func(r rune, i, j int) {
		if s.At(i, j) == GALAXY {
			galaxies = append(galaxies, Coordinate{j, i})
		}
	})
	return galaxies
}

func Distance(from, to Coordinate) int {
	return int(math.Abs(float64(to.x-from.x))) + int(math.Abs(float64(to.y-from.y)))
}

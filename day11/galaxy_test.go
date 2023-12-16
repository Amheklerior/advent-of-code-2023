package day11

import (
	"fmt"
	"math"
	"math/rand"
	"slices"
	"testing"

	"amheklerior.com/advent-of-code-2023/utils"
)

func TestEmptyRow(t *testing.T) {
	expected := []rune{EMPTY, EMPTY, EMPTY}
	created := EmptyRow(3)

	if !slices.Equal(created, expected) {
		t.Errorf("Got %v, expected %v", created, expected)
	}
}

func TestGalaxiesMap(t *testing.T) {
	space := Space{
		{EMPTY, GALAXY, EMPTY},
		{EMPTY, EMPTY, GALAXY},
		{GALAXY, EMPTY, EMPTY},
	}
	expected := []Coordinate{
		{1, 0},
		{2, 1},
		{0, 2},
	}

	found := space.GalaxiesMap()

	if !slices.Equal(found, expected) {
		t.Errorf("Got %v, expected %v", found, expected)
	}
}

func TestDistance(t *testing.T) {
	for i := 0; i < 25; i++ {
		c1 := Coordinate{rand.Intn(50), rand.Intn(50)}
		c2 := Coordinate{rand.Intn(50), rand.Intn(50)}
		t.Run(fmt.Sprintf("Test distance between %v and %v", c1, c2), func(t *testing.T) {
			expected := int(
				math.Abs((float64(c2.x - c1.x))) +
					math.Abs(float64(c2.y-c1.y)))
			d := Distance(c1, c2)
			if d != int(expected) {
				t.Errorf("Got %v, expected %v. For coordinates %v and %v", d, expected, c1, c2)
			}
		})
	}
}

func TestExpansion(t *testing.T) {
	space := Space{
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, GALAXY, GALAXY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, GALAXY},
	}

	expected := Space{
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, GALAXY, GALAXY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
		{EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, GALAXY},
	}

	expanded := space.Expand()

	for i, row := range expanded {
		if !slices.Equal(row, expected[i]) {
			t.Errorf("Got %v, Expected %v, on row #%v", row, expected[i], i)
		}
	}
}

func TestManual(t *testing.T) {
	space := NewSpace(utils.ReadFile("./data/p1-input.test.txt"))
	space.Expand()
	galaxies := space.GalaxiesMap()
	fmt.Printf("Galaxies 5 and 9 are %v light years apart from each others\n", Distance(galaxies[4], galaxies[8]))
	fmt.Printf("Galaxies 1 and 7 are %v light years apart from each others\n", Distance(galaxies[0], galaxies[6]))
	fmt.Printf("Galaxies 3 and 6 are %v light years apart from each others\n", Distance(galaxies[2], galaxies[5]))
	fmt.Printf("Galaxies 8 and 9 are %v light years apart from each others\n", Distance(galaxies[7], galaxies[8]))
}

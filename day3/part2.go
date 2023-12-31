package day3

import (
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	scanner := utils.Scanner(content)

	lineIdx := 0
	var matrix [][]rune
	var parts []Part
	var gears []Gear
	for scanner.Scan() {
		line := scanner.Text()

		// Build the matrix
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)

		// Collect the parts' info
		numbers := utils.GetOccurrences(line, `\d+`)
		for _, n := range numbers {
			idx := strings.Index(line, n)
			parts = append(parts, Part{
				number: n,
				pos:    [2]int{lineIdx, idx},
			})
			str := "."
			if len(n) == 2 {
				str = ".."
			}
			if len(n) == 3 {
				str = "..."
			}
			line = strings.Replace(line, n, str, 1)
		}

		// Collect the gears info
		asterics := utils.GetOccurrences(line, `\*`)
		for _, a := range asterics {
			idx := strings.Index(line, a)
			gears = append(gears, Gear{lineIdx, idx})
			line = strings.Replace(line, a, ".", 1)
		}

		lineIdx++
	}

	// Compute gears' ratios and sum them up
	sum := 0
	for _, gear := range gears {
		var adjacents []int
		for _, part := range parts {
			verticallyWithinBoundary := part.pos[0] >= gear.x-1 && part.pos[0] <= gear.x+1
			horizontallyWithinBoundary := part.pos[1] >= gear.y-len(part.number) && part.pos[1] <= gear.y+1
			if verticallyWithinBoundary && horizontallyWithinBoundary {
				adjacents = append(adjacents, utils.ToInt(part.number))
			}
		}
		if len(adjacents) == 2 {
			sum += adjacents[0] * adjacents[1]
		}
	}
	return sum
}

func TestP2() {
	utils.Run(3, 2, 467835, SolutionPart2)
}

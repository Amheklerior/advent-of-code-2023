package day3

import (
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

func TODO1(str string) int {
	return len(str)
}

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	scanner := utils.Scanner(content)

	lineIdx := 0
	var matrix [][]rune
	var parts []Part
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

		lineIdx++
	}

	// Check valid parts and sum them up
	sum := 0
	for _, p := range parts {
		isValid := p.isValid(matrix)
		if isValid {
			sum += utils.ToInt(p.number)
		}
	}
	return sum
}

func TestP1() {
	utils.Run(3, 1, 4361, SolutionPart1)
}

package day0

import (
	"amheklerior.com/advent-of-code-2023/utils"
)

func TODO1(str string) int {
	return len(str)
}

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	scanner := utils.Scanner(content)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += TODO1(line)
	}
	return sum
}

func TestP1() {
	utils.Run(0, 1, -1, SolutionPart1)
}

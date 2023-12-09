package day0

import (
	"amheklerior.com/advent-of-code-2023/utils"
)

func TODO2(str string) int {
	return len(str)
}

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	scanner := utils.Scanner(content)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += TODO2(line)
	}
	return sum
}

func TestP2() {
	utils.Run(0, 2, -1, SolutionPart2)
}

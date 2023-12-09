package day1

import (
	"fmt"

	"amheklerior.com/advent-of-code-2023/utils"
)

func getCalibrationValueP1(line string) int {
	// extract all digits in the line and take first and last
	numbers, found := findDigits(line)
	if !found {
		return 0
	}

	first := getFirst(numbers)
	last := getLast(numbers)

	// Concatenate the two digits found to create the whole line value
	str_value := fmt.Sprintf("%s%s", first, last)
	return utils.ToInt(str_value)
}

func SolutionPart1(path string) int {
	content := utils.ReadFile(path)
	scanner := utils.Scanner(content)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += getCalibrationValueP1(line)
	}
	return sum
}

func TestP1() {
	utils.Run(1, 1, 142, SolutionPart1)
}

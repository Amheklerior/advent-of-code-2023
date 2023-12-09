package day1

import (
	"fmt"

	"amheklerior.com/advent-of-code-2023/utils"
)

func getCalibrationValueP2(line string) int {
	// extract all digits in the line and take first and last
	numbers, found := findDigits(line)
	var first, last string
	if found {
		first = getFirst(numbers)
		last = getLast(numbers)
	}

	// Check whether there was a full-spelled digit (like "one", "seven", etc)
	// before the first, and use that as first if found
	prefix, found := getPrefix(line)
	if found {
		prefix = replaceWithDigits(prefix, false)
		nums, notEmpty := findDigits(prefix)
		if notEmpty {
			first = getFirst(nums)
		}
	}

	// Check wheter there was a non-digit number after the last digit found
	// If there was one, use that as the last number
	suffix, found := getSuffix(line)
	if found {
		suffix = replaceWithDigits(suffix, true)
		nums, notEmpty := findDigits(suffix)
		if notEmpty {
			last = getLast(nums)
		}
	}

	// Concatenate the two digits found to create the whole line value
	str_value := fmt.Sprintf("%s%s", first, last)
	return utils.ToInt(str_value)
}

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	scanner := utils.Scanner(content)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += getCalibrationValueP2(line)
	}
	return sum
}

func TestP2() {
	utils.Run(1, 2, 281, SolutionPart2)
}

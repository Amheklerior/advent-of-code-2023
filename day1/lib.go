package day1

import (
	"regexp"
	"strings"
)

func getFirst(nums []string) string {
	return nums[0]
}

func getLast(nums []string) string {
	return nums[len(nums)-1]
}

func findDigits(str string) ([]string, bool) {
	numbers := regexp.MustCompile(`\d`).FindAllString(str, -1)
	return numbers, len(numbers) > 0
}

func getPrefix(str string) (string, bool) {
	prefix := regexp.MustCompile(`^\D+`).FindString(str)
	return prefix, len(prefix) > 0
}

func getSuffix(str string) (string, bool) {
	suffix := regexp.MustCompile(`\D+$`).FindString(str)
	return suffix, len(suffix) > 0
}

func replaceWithDigits(str string, preserveLast bool) string {
	replacer := strings.NewReplacer(
		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	if preserveLast {
		replacer = strings.NewReplacer(
			"oneight", "8",
			"twone", "1",
			"threeight", "8",
			"fiveight", "8",
			"sevenine", "9",
			"eightwo", "2",
			"eighthree", "3",
			"nineight", "8",
			"one", "1",
			"two", "2",
			"three", "3",
			"four", "4",
			"five", "5",
			"six", "6",
			"seven", "7",
			"eight", "8",
			"nine", "9",
		)
	}
	return replacer.Replace(str)
}

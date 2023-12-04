package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const PATH string = "./input.txt"

func getFirst(nums []string) string {
	return nums[0]
}

func getLast(nums []string) string {
	return nums[len(nums)-1]
}

func main() {
	f, e := os.Open(PATH)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	digitsRegex := regexp.MustCompile(`\d`)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// extract all digits in the line and take first and last
		numbers := digitsRegex.FindAllString(line, -1)
		first := getFirst(numbers)
		last := getLast(numbers)

		// Check whether there was a full-spelled digit (like "one", "seven", etc)
		// before the first, and use that as first if found
		prefixRegex := regexp.MustCompile(`^\D+`)
		prefixMatch := prefixRegex.FindString(line)
		if len(prefixMatch) > 0 {
			prefixReplacer := strings.NewReplacer("one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
			prefix := prefixReplacer.Replace(prefixMatch)
			prefixNums := digitsRegex.FindAllString(prefix, -1)
			if len(prefixNums) > 0 {
				first = getFirst(prefixNums)
			}
		}

		// Check wheter there was a non-digit number after the last digit found
		// If there was one, use that as the last number
		suffixRegex := regexp.MustCompile(`\D+$`)
		suffixMatch := suffixRegex.FindString(line)
		if len(suffixMatch) > 0 {
			suffixReplacer := strings.NewReplacer("oneight", "8", "twone", "1", "threeight", "8", "fiveight", "8", "sevenine", "9", "eightwo", "2", "eighthree", "3", "nineight", "8", "one", "1", "two", "2", "three", "3", "four", "4", "five", "5", "six", "6", "seven", "7", "eight", "8", "nine", "9")
			suffix := suffixReplacer.Replace(suffixMatch)
			suffixNums := digitsRegex.FindAllString(suffix, -1)
			if len(suffixNums) > 0 {
				last = getLast(suffixNums)
			}
		}

		// Concatenate the two digits found to create the whole line value
		line_value := fmt.Sprintf("%s%s", first, last)
		calibration_value, _ := strconv.Atoi(line_value)

		// add the value found in the current line to the global sum for the entire file
		sum += int(calibration_value)
	}

	fmt.Println(sum)
}

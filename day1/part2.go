package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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
	value, _ := strconv.Atoi(str_value)
	return value
}

func SolutionPart2(path string) int {
	f, e := os.Open(path)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += getCalibrationValueP2(line)
	}
	return sum
}

func TestP2() {
	fmt.Println("Day 1 / Part 2: Test")
	expected := 281
	result := SolutionPart2("./day1/data/p2-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

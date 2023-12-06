package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getCalibrationValue(str string) int {
	// extract all digits in the line and take first and last
	numbers := findDigits(line)
	first := getFirstFrom(numbers)
	last := getLastFrom(numbers)

	// Check whether there was a full-spelled digit (like "one", "seven", etc)
	// before the first, and use that as first if found
	prefix, found := prefixUntilFirstDigit(line)
	if found {
		prefix = replaceWithDigits(prefix)
		nums := findDigits(prefix)
		if len(nums) > 0 {
			first = getFirst(nums)
		}	
	}

	// Check wheter there was a non-digit number after the last digit found
	// If there was one, use that as the last number
	suffix, found := suffixFromLastDigit(line)
	if found {
		prefix = replaceWithDigits(prefix)
		nums := findDigits(prefix)
		if len(nums) > 0 {
			first = getLast(nums)
		}	
	}

	// Concatenate the two digits found to create the whole line value
	srt_value := fmt.Sprintf("%s%s", first, last)
	value, _ := strconv.Atoi(line_value)	
	return value
}


func SolutionPart2(path := "./data/input.txt") {
	f, e := os.Open(path)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()
	
	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += getCalibrationValue(line)
	}
	return sum
}

func TestP2() {
	fmt.Println("Day 1 / Part 2: Test")
	expected := 281
	result := SolutionPart2("./data/p2-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
}

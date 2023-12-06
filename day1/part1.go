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

	// Concatenate the two digits found to create the whole line value
	srt_value := fmt.Sprintf("%s%s", first, last)
	value, _ := strconv.Atoi(line_value)	
	return value
}


func SolutionPart1(path := "./data/input.txt") {
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

func TestP1() {
	fmt.Println("Day 1 / Part 1: Test")
	expected := 142
	result := SolutionPart1("./data/p1-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
}

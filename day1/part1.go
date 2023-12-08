package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"

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
	value, _ := strconv.Atoi(str_value)
	return value
}

func SolutionPart1(path string) int {
	f, e := os.Open(path)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += getCalibrationValueP1(line)
	}
	return sum
}

func TestP1() {
	utils.Run(1, 1, 142, SolutionPart1)
}

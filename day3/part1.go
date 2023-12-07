package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func TODO1(str string) int {
	return len(str)
}

func SolutionPart1(path string) int {
	f, e := os.Open(path)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	r := regexp.MustCompile(`\d+`)
	scanner := bufio.NewScanner(f)

	lineIdx := 0
	var matrix [][]rune
	var parts []Part
	for scanner.Scan() {
		line := scanner.Text()

		// Build the matrix
		var row []rune
		for _, char := range line {
			row = append(row, char)
		}
		matrix = append(matrix, row)

		// Collect the parts' info
		numbers := r.FindAllString(line, -1)
		for _, n := range numbers {
			idx := strings.Index(line, n)
			parts = append(parts, Part{
				number: n,
				pos:    [2]int{lineIdx, idx},
			})
			str := "."
			if len(n) == 2 {
				str = ".."
			}
			if len(n) == 3 {
				str = "..."
			}
			line = strings.Replace(line, n, str, 1)
		}

		lineIdx++
	}

	// Check valid parts and sum them up
	sum := 0
	for _, p := range parts {
		isValid := p.isValid(matrix)
		if isValid {
			partId, _ := strconv.Atoi(p.number)
			sum += partId
		}
	}
	return sum
}

func TestP1() {
	fmt.Println("Day 3 / Part 1: Test")
	expected := 4361
	result := SolutionPart1("./day3/data/p1-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}
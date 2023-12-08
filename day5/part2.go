package day5

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func TODO2(str string) int {
	return len(str)
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
		sum += TODO2(line)
	}
	return sum
}

func TestP2() {
	fmt.Println("Day 5 / Part 2: Test")
	expected := 142
	result := SolutionPart2("./day5/data/p2-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

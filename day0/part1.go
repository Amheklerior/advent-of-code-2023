package day0

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	sum := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += TODO1(line)
	}
	return sum
}

func TestP1() {
	fmt.Println("Day 0 / Part 1: Test")
	expected := 142
	result := SolutionPart1("./day0/data/p1-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

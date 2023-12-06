package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func findMax(counts []string) int {
	sort.Slice(counts, func(i, j int) bool {
		n1, _ := strconv.Atoi(counts[i])
		n2, _ := strconv.Atoi(counts[j])
		return n1 < n2
	})
	max, _ := strconv.Atoi(counts[len(counts)-1])
	return max
}

func fewestNumberOfCubes(game string) (int, int, int) {
	// Extract the counts for all colored cubes
	reds := getCountsFor(game, "red")
	greens := getCountsFor(game, "green")
	blues := getCountsFor(game, "blue")

	// Find the highest count for each colored cube
	maxRed := findMax(reds)
	maxGreen := findMax(greens)
	maxBlue := findMax(blues)

	return maxRed, maxGreen, maxBlue
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
		red, green, blue := fewestNumberOfCubes(line)
		sum += red * green * blue
	}
	return sum
}

func TestP2() {
	fmt.Println("Day 2 / Part 2: Test")
	expected := 2286
	result := SolutionPart2("./day2/data/p2-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

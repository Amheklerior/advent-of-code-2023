package day5

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func extractSeeds(line string) []int {
	var seeds []int
	prefix := regexp.MustCompile(`seeds:`).FindString(line)
	line = strings.TrimPrefix(line, prefix)
	ids := strings.Fields(line)
	for _, id := range ids {
		seed, _ := strconv.Atoi(id)
		seeds = append(seeds, seed)
	}
	return seeds
}

func buildDataStructures(input string) ([]int, [][][]int) {
	var seeds []int
	var pipeline [][][]int

	mapId := -1
	scanner := Scanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		// skip empty lines
		if line == "" {
			continue
		}

		// if its the seeds line extract all seeds
		if strings.Contains(line, "seeds") {
			seeds = extractSeeds(line)
			continue
		}

		prefix := regexp.MustCompile(`.+:`).FindString(line)

		// if its the heading of a new map, create a new one in the pipeline
		if strings.Contains(prefix, "map") {
			pipeline = append(pipeline, make([][]int, 0, 10))
			mapId++
			continue
		}

		// it's a map instruction line
		var values []int
		data := strings.Fields(line)
		for _, x := range data {
			v, _ := strconv.Atoi(x)
			values = append(values, v)
		}
		pipeline[mapId] = append(pipeline[mapId], values)
	}

	return seeds, pipeline
}

type Range struct {
	start, end int
}

func (r *Range) contains(n int) bool {
	return n >= r.start && n < r.end
}

func extractData(m []int) (Range, int) {
	destinationStart, sourceStart, rangeLenght := m[0], m[1], m[2]
	rangeStart := sourceStart
	rangeEnd := sourceStart + rangeLenght
	gap := destinationStart - sourceStart

	return Range{rangeStart, rangeEnd}, gap
}

func getDestination(source int, mapper [][]int) int {
	for _, v := range mapper {
		sourceRange, gap := extractData(v)
		if sourceRange.contains(source) {
			return source + gap
		}
	}
	return source
}

func ReadFile(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Could not open the file: %s", err)
	}
	content := string(bytes)
	return content
}

func Scanner(fileContent string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(fileContent))
}

func passThroughPipeline(seed int, pipeline [][][]int) int {
	pointer := seed
	for _, sourceToDestinationMap := range pipeline {
		pointer = getDestination(pointer, sourceToDestinationMap)
	}
	return pointer
}

func SolutionPart1(path string) int {
	content := ReadFile(path)

	seeds, pipeline := buildDataStructures(content)

	minLocation := math.MaxInt

	for _, seed := range seeds {
		location := passThroughPipeline(seed, pipeline)
		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func TestP1() {
	fmt.Println("Day 5 / Part 1: Test")
	expected := 35
	result := SolutionPart1("./day5/data/p1-input.test.txt")

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
	}
	println()
}

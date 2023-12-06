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

type Part struct {
	number string
	pos    [2]int
}

type Gear struct {
	x, y int
}

func main() {
	f, e := os.Open(PATH)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	r := regexp.MustCompile(`\d+`)
	s := regexp.MustCompile(`\*`)
	scanner := bufio.NewScanner(f)

	lineIdx := 0
	var matrix [][]rune
	var parts []Part
	var gears []Gear
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

		// Collect the gears info
		asterics := s.FindAllString(line, -1)
		for _, a := range asterics {
			idx := strings.Index(line, a)
			gears = append(gears, Gear{lineIdx, idx})
			line = strings.Replace(line, a, ".", 1)
		}

		lineIdx++
	}

	// Compute gears' ratios and sum them up
	sum := 0
	for _, gear := range gears {
		var adjacents []int
		for _, part := range parts {
			verticallyWithinBoundary := part.pos[0] >= gear.x-1 && part.pos[0] <= gear.x+1
			horizontallyWithinBoundary := part.pos[1] >= gear.y-len(part.number) && part.pos[1] <= gear.y+1
			if verticallyWithinBoundary && horizontallyWithinBoundary {
				num, _ := strconv.Atoi(part.number)
				adjacents = append(adjacents, num)
			}
		}
		if len(adjacents) == 2 {
			sum += adjacents[0] * adjacents[1]
		}
	}

	fmt.Println(sum)
}

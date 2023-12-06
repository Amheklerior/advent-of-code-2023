package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const PATH string = "./input.txt"

type Part struct {
	number string
	pos    [2]int
}

func isSpecialChar(char rune) bool {
	return char != '.' && !unicode.IsDigit(char)
}

func (p *Part) isValid(mtx [][]rune) bool {
	for i, _ := range p.number {
		x, y := p.pos[0], p.pos[1]+i

		onTopRow := x == 0
		onBottomRow := x == len(mtx)-1
		onFirstCol := y == 0
		onLastCol := y == len(mtx[x])-1
		onLeftmostDigit := i == 0
		onRightmostDigit := i == len(p.number)-1

		// check NORD
		if onLeftmostDigit && !onTopRow && isSpecialChar(mtx[x-1][y]) {
			return true
		}
		// check SOUTH
		if onLeftmostDigit && !onBottomRow && isSpecialChar(mtx[x+1][y]) {
			return true
		}
		// ckeck WEST
		if onLeftmostDigit && !onFirstCol && isSpecialChar(mtx[x][y-1]) {
			return true
		}
		// check EAST
		if onRightmostDigit && !onLastCol && isSpecialChar(mtx[x][y+1]) {
			return true
		}
		// check NW
		if onLeftmostDigit && !onFirstCol && !onTopRow && isSpecialChar(mtx[x-1][y-1]) {
			return true
		}
		// check NE
		if !onLastCol && !onTopRow && isSpecialChar(mtx[x-1][y+1]) {
			return true
		}
		// check SW
		if onLeftmostDigit && !onFirstCol && !onBottomRow && isSpecialChar(mtx[x+1][y-1]) {
			return true
		}
		// check SE
		if !onLastCol && !onBottomRow && isSpecialChar(mtx[x+1][y+1]) {
			return true
		}
	}

	return false
}

func main() {
	f, e := os.Open(PATH)
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
	fmt.Println(sum)
}

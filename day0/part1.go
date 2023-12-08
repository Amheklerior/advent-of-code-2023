package day0

import (
	"bufio"
	"log"
	"os"

	"amheklerior.com/advent-of-code-2023/utils"
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
	utils.Run(0, 1, -1, SolutionPart1)
}

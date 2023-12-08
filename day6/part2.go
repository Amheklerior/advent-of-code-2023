package day6

import (
	"bufio"
	"log"
	"os"

	"amheklerior.com/advent-of-code-2023/utils"
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
	utils.Run(6, 1, -1, SolutionPart2)
}

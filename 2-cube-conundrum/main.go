package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const PATH string = "./input.txt"

func getCountsFor(str, color string) []string {
	digitsRegex := regexp.MustCompile(`\d+`)
	r := regexp.MustCompile(fmt.Sprintf(`\d+ %s`, color))
	return digitsRegex.FindAllString(strings.Join(r.FindAllString(str, -1), " "), -1)
}

func findMax(counts []string) int {
	sort.Slice(counts, func(i, j int) bool {
		n1, _ := strconv.Atoi(counts[i])
		n2, _ := strconv.Atoi(counts[j])
		return n1 < n2
	})
	max, _ := strconv.Atoi(counts[len(counts)-1])
	return max
}

func main() {
	f, e := os.Open(PATH)
	if e != nil {
		log.Fatalf("Could not open the file: %s", e)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		// Extract the counts for all colored cubes
		reds := getCountsFor(line, "red")
		greens := getCountsFor(line, "green")
		blues := getCountsFor(line, "blue")

		// Find the highest count for each colored cube
		maxRed := findMax(reds)
		maxGreen := findMax(greens)
		maxBlue := findMax(blues)

		// Sum up the power of the set of cubes
		sum += maxRed * maxGreen * maxBlue
	}

	fmt.Println(sum)
}

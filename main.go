package main

import (
	"fmt"

	"amheklerior.com/advent-of-code-2023/day6"
)

func input(day int) string {
	return fmt.Sprintf("./day%v/data/input.txt", day)
}

func main() {
	fmt.Println("************ Advent Of Code 2023 ************")
	println()

	// day1.TestP1()
	// day1.TestP2()
	// day2.TestP1()
	// day2.TestP2()
	// day3.TestP1()
	// day3.TestP2()
	// day4.TestP1()
	// day4.TestP2()
	// day5.TestP1()
	// day5.TestP2()
	// day6.TestP1()
	// day6.TestP2()

	println()

	/* TODO: write a CLI which takes day and part and a flag for testing and runs only the proper code
	fmt.Printf("Day 1 / Part 1 --> %v", day1.SolutionPart1(input(1)))
	fmt.Println()
	fmt.Printf("Day 1 / Part 2 --> %v", day1.SolutionPart2(input(1)))
	fmt.Println()

	fmt.Printf("Day 2 / Part 1 --> %v", day2.SolutionPart1(input(2)))
	fmt.Println()
	fmt.Printf("Day 2 / Part 2 --> %v", day2.SolutionPart2(input(2)))
	fmt.Println()

	fmt.Printf("Day 3 / Part 1 --> %v", day3.SolutionPart1(input(3)))
	fmt.Println()
	fmt.Printf("Day 3 / Part 2 --> %v", day3.SolutionPart2(input(3)))
	fmt.Println()

	fmt.Printf("Day 4 / Part 1 --> %v", day4.SolutionPart1(input(4)))
	fmt.Println()
	fmt.Printf("Day 4 / Part 2 --> %v", day4.SolutionPart2(input(4)))
	fmt.Println()

	fmt.Printf("Day 5 / Part 1 --> %v", day5.SolutionPart1(input(5)))
	fmt.Println()
	fmt.Printf("Day 5 / Part 2 --> %v", day5.SolutionPart2(input(5)))
	fmt.Println()

	*/
	fmt.Printf("Day 6 / Part 1 --> %v", day6.SolutionPart1(input(6)))
	fmt.Println()
	fmt.Printf("Day 6 / Part 2 --> %v", day6.SolutionPart2(input(6)))
	fmt.Println()

}

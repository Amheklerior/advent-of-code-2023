package utils

import "fmt"

func Run(day, part, expected int, computeSolution func(string) int) bool {
	fmt.Printf("Day %v / Part %v: Test", day, part)
	fmt.Println()

	inputPath := fmt.Sprintf("./day%v/data/p%v-input.test.txt", day, part)
	result := computeSolution(inputPath)

	if result == expected {
		fmt.Printf("Success!! Result is: %v", result)
		fmt.Println()
		return true
	} else {
		fmt.Printf("Failure! Expected %v, got %v", expected, result)
		fmt.Println()
		return false
	}
}

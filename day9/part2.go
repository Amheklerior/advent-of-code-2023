package day9

import (
	"sync"

	"amheklerior.com/advent-of-code-2023/day9/oasis"
	"amheklerior.com/advent-of-code-2023/utils"
)

func SolutionPart2(path string) int {
	content := utils.ReadFile(path)
	sequences := oasis.ParseInput(content)
	var wg sync.WaitGroup
	sum := 0
	for _, seq := range sequences {
		wg.Add(1)
		go func(sequence oasis.Sequence) {
			defer wg.Done()
			sum += sequence.Previous()
		}(seq)
	}
	wg.Wait()
	return sum
}

func TestP2() {
	utils.Run(9, 2, 2, SolutionPart2)
}

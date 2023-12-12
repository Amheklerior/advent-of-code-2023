package oasis

import "amheklerior.com/advent-of-code-2023/utils"

func ParseInput(content string) []Sequence {
	scanner := utils.Scanner(content)
	var sequences []Sequence
	for scanner.Scan() {
		line := scanner.Text()
		sequences = append(sequences, NewSequence(line))
	}
	return sequences
}

package oasis

import (
	"fmt"
	"slices"
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

type Sequence []int

func NewSequence(input string) Sequence {
	var seq Sequence
	fields := strings.Fields(input)
	for _, v := range fields {
		seq = append(seq, utils.ToInt(v))
	}
	return seq
}

func (s *Sequence) Next() int {
	reduced := *s
	reductionStack := []Sequence{reduced}
	fmt.Printf("%v\n", reduced)
	for !reduced.IsZeroSequence() {
		reduced = reduced.Reduce()
		fmt.Printf("%v\n", reduced)
		reductionStack = slices.Insert(reductionStack, 0, reduced)
	}
	next := 0
	for _, seq := range reductionStack {
		next += seq[len(seq)-1]
	}
	fmt.Printf("Next: %v\n\n", next)
	return next
}

func (s *Sequence) Previous() int {
	reduced := *s
	reductionStack := []Sequence{reduced}
	fmt.Printf("%v\n", reduced)
	for !reduced.IsZeroSequence() {
		reduced = reduced.Reduce()
		fmt.Printf("%v\n", reduced)
		reductionStack = slices.Insert(reductionStack, 0, reduced)
	}
	previous := 0
	for _, seq := range reductionStack {
		previous = seq[0] - previous
	}
	fmt.Printf("Previous: %v\n\n", previous)
	return previous
}

func (s *Sequence) Reduce() Sequence {
	this := *s
	length := len(this)

	if length < 2 {
		return this
	}

	if length == 2 {
		return Sequence{this[1] - this[0]}
	}

	leftHalf := this[0 : length/2+(length%2)]
	rightHalf := this[length/2:]
	reducedLeft := leftHalf.Reduce()
	reducedRight := rightHalf.Reduce()

	reduced := make(Sequence, 0)
	reduced = append(reduced, reducedLeft...)
	if length%2 == 0 {
		reduced = append(reduced, this[length/2]-this[length/2-1])
	}
	reduced = append(reduced, reducedRight...)

	return reduced
}

func (s *Sequence) IsZeroSequence() bool {
	for _, num := range *s {
		if num != 0 {
			return false
		}
	}
	return true
}

func (s *Sequence) Equal(other Sequence) bool {
	for i, val := range *s {
		if val != other[i] {
			return false
		}
	}
	return true
}

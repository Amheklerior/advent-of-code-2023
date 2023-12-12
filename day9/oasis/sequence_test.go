package oasis

import (
	"testing"
)

func TestNext(t *testing.T) {
	input := Sequence{1, 3, 6, 10, 15, 21}
	expected := 28

	result := input.Next()

	if result != expected {
		t.Errorf("Error calculating next item in sequence %v. Expected %v, got %v instead", input, expected, result)
	}
}

func TestPrevious(t *testing.T) {
	input := Sequence{10, 13, 16, 21, 30, 45}
	expected := 5

	result := input.Previous()

	if result != expected {
		t.Errorf("Error calculating previous item in sequence %v. Expected %v, got %v instead", input, expected, result)
	}
}

func TestReduce(t *testing.T) {
	input := Sequence{1, 3, 6, 8}
	expected := Sequence{2, 3, 2}

	result := input.Reduce()

	if !result.Equal(expected) {
		t.Errorf("Error reducing sequence %v. Expected %v, got %v instead", input, expected, result)
	}
}

func TestIsZeroSequence(t *testing.T) {
	nonZeroSequence := Sequence{1, 2, 3, 4, 5}
	zeroSequence := Sequence{0, 0, 0, 0, 0, 0, 0}

	if !zeroSequence.IsZeroSequence() {
		t.Errorf("Error for sequence %v. Expected to be valuated as a zero sequence", zeroSequence)
	}
	if nonZeroSequence.IsZeroSequence() {
		t.Errorf("Error for sequence %v. Expected to NOT be valuated as a zero sequence", nonZeroSequence)
	}
}

func TestSequenceEquality(t *testing.T) {
	seqA := Sequence{1, 2, 3, 4, 5}
	seqB := Sequence{1, 2, 3, 4, 5}
	seqC := Sequence{5, 3, 6, 9, 2}

	if !seqA.Equal(seqB) {
		t.Errorf("Expected sequences A:%v and B:%v to equal.", seqA, seqB)
	}
	if seqA.Equal(seqC) {
		t.Errorf("Expected sequences A:%v and C:%v to differs.", seqA, seqC)
	}
}

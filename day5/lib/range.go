package lib

type Range struct {
	Start, End int
}

func NewRange(start, end int) Range {
	return Range{start, end}
}

func (r *Range) Contains(n int) bool {
	return n >= r.Start && n < r.End
}

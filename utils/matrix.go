package utils

type Matrix[T any] [][]T

func (m *Matrix[T]) At(i, j int) T {
	return (*m)[i][j]
}

func (m *Matrix[T]) Height() int {
	return len(*m)
}

func (m *Matrix[T]) Width() int {
	if m.Height() == 0 {
		return 0
	}
	return len((*m)[0])
}

func (m *Matrix[T]) ForEach(callback func(T, int, int)) {
	for i := range *m {
		for j := range (*m)[i] {
			callback((*m)[i][j], i, j)
		}
	}
}

package utils

func GreatestCommonDivisor(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LeastCommonMultiple(a, b int) int {
	return (a * b) / GreatestCommonDivisor(a, b)
}

type IntSlice []int

func (numbers IntSlice) LeastCommonMultiple() int {
	result := 1
	for _, num := range numbers {
		result = LeastCommonMultiple(result, num)
	}
	return result
}

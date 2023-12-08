package lib

func extractData(m []int) (Range, int) {
	destinationStart, sourceStart, rangeLenght := m[0], m[1], m[2]
	rangeStart := sourceStart
	rangeEnd := sourceStart + rangeLenght
	gap := destinationStart - sourceStart

	return NewRange(rangeStart, rangeEnd), gap
}

func getDestination(source int, mapper [][]int) int {
	for _, v := range mapper {
		sourceRange, gap := extractData(v)
		if sourceRange.Contains(source) {
			return source + gap
		}
	}
	return source
}

func PassThroughPipeline(seed int, pipeline [][][]int) int {
	pointer := seed
	for _, sourceToDestinationMap := range pipeline {
		pointer = getDestination(pointer, sourceToDestinationMap)
	}
	return pointer
}

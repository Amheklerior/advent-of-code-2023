package lib

import (
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

func BuildDataStructures(input string) [][][]int {
	var pipeline [][][]int

	mapId := -1
	scanner := utils.Scanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		// skip seeds line and empty lines
		if line == "" || strings.Contains(line, "seeds") {
			continue
		}

		_, prefix := utils.ExtractPrefix(line, `.+:`)

		// if its the heading of a new map, create a new one in the pipeline
		if strings.Contains(prefix, "map") {
			pipeline = append(pipeline, make([][]int, 0, 10))
			mapId++
			continue
		}

		// it's a map instruction line
		var values []int
		data := strings.Fields(line)
		for _, v := range data {
			values = append(values, utils.ToInt(v))
		}
		pipeline[mapId] = append(pipeline[mapId], values)
	}

	return pipeline
}

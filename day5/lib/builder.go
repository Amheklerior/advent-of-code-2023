package lib

import (
	"regexp"
	"strconv"
	"strings"
)

func BuildDataStructures(input string) [][][]int {
	var pipeline [][][]int

	mapId := -1
	scanner := Scanner(input)
	for scanner.Scan() {
		line := scanner.Text()

		// skip seeds line and empty lines
		if line == "" || strings.Contains(line, "seeds") {
			continue
		}

		prefix := regexp.MustCompile(`.+:`).FindString(line)

		// if its the heading of a new map, create a new one in the pipeline
		if strings.Contains(prefix, "map") {
			pipeline = append(pipeline, make([][]int, 0, 10))
			mapId++
			continue
		}

		// it's a map instruction line
		var values []int
		data := strings.Fields(line)
		for _, x := range data {
			v, _ := strconv.Atoi(x)
			values = append(values, v)
		}
		pipeline[mapId] = append(pipeline[mapId], values)
	}

	return pipeline
}

package navigator

import (
	"strings"

	"amheklerior.com/advent-of-code-2023/utils"
)

func removeUselessRunes(str string) string {
	replacer := strings.NewReplacer("(", "", ")", "", " ", "")
	return replacer.Replace(str)
}

func NewNetwork(input string) (Network, []Node) {
	var entryPoints []Node
	network := make(Network)

	scanner := utils.Scanner(input)
	for scanner.Scan() {
		line := removeUselessRunes(scanner.Text())

		if line == "" || !strings.Contains(line, "=") {
			continue
		}

		parts := strings.Split(line, "=")
		node := Node(strings.TrimSpace(parts[0]))
		branches := strings.Split(parts[1], ",")

		network[node] = make(Brances)
		network[node][LEFT] = Node(branches[0])
		network[node][RIGHT] = Node(branches[1])

		if node.IsEntryPoint() {
			entryPoints = append(entryPoints, node)
		}
	}

	return network, entryPoints
}

func NewInstructionPath(input string) Path {
	scanner := utils.Scanner(input)
	scanner.Scan()
	return Path(strings.TrimSpace(scanner.Text()))
}

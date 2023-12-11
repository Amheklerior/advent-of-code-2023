package navigator

type Network map[Node]Brances
type Brances map[Direction]Node
type Path []Direction

type Navigator struct {
	network      Network
	instructions Path
	entryPoints  []Node
}

func (network *Network) Traverse(instructions Path) int {
	current := START
	count := 0
	for i := 0; !current.IsExitPoint(); i = (i + 1) % len(instructions) {
		direction := instructions[i]
		current = (*network)[current][direction]
		count++
	}
	return count
}

func (network *Network) FindPathLenghtFrom(entryPoint Node, instructions Path) int {
	direction := instructions[0]
	current := (*network)[entryPoint][direction]
	hops := 1
	for i := 1; !current.IsExitPoint(); i = (i + 1) % len(instructions) {
		direction = instructions[i]
		current = (*network)[current][direction]
		hops++
	}
	return hops
}

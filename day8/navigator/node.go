package navigator

type Node string

const (
	START       Node = "AAA"
	DESTINATION Node = "ZZZ"
)

func (node *Node) IsEntryPoint() bool {
	n := string(*node)
	return n[len(n)-1] == 'A'
}

func (node *Node) IsExitPoint() bool {
	n := string(*node)
	return n[len(n)-1] == 'Z'
}

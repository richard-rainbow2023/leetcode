package letture03

// Node
/* Definition for a Node. */
type Node struct {
	Val      int
	Children []*Node
}

var ansN []int

func preorder(root *Node) []int {
	ansN = []int{}
	dfsN(root)
	return ansN
}

func dfsN(root *Node) {
	if root == nil {
		return
	}
	ansN = append(ansN, root.Val)
	for _, child := range root.Children {
		dfsN(child)
	}
}

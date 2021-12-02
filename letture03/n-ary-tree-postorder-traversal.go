package letture03

var ansP []int

func postorder(root *Node) []int {
	ansP = []int{}
	dfsP(root)
	return ansP
}

func dfsP(root *Node) {
	if root == nil {
		return
	}
	for _, child := range root.Children {
		dfsN(child)
	}
	ansP = append(ansP, root.Val)
}

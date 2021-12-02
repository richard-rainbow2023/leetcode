package letture03

func maxDepth(root *Node) int {
	maxChildDepth := 0
	if root == nil {
		return 0
	}
	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > maxChildDepth {
			maxChildDepth = childDepth
		}
	}
	return maxChildDepth + 1
}

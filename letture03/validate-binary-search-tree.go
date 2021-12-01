package letture03

import "math"

type BTreeNode struct {
	Val   int
	Left  *BTreeNode
	Right *BTreeNode
}

func isValidBST(root *BTreeNode) bool {
	return cal(root, math.MinInt64, math.MaxInt64)
}

func cal(root *BTreeNode, leftRange int64, rightRange int64) bool {
	if root == nil {
		return true
	}
	if int64(root.Val) > rightRange || int64(root.Val) < leftRange {
		return false
	}
	return cal(root.Left, leftRange, int64(root.Val)-1) &&
		cal(root.Right, int64(root.Val)-1, rightRange)
}

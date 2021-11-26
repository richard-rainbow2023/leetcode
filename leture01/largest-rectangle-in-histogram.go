package leture01

import "container/list"

type Rect struct {
	width  int
	height int
}

func largestRectangleArea(heights []int) int {
	stack := list.New()
	heights = append(heights, 0)
	ans := 0
	for _, height := range heights {
		accumulateWidth := 0
		for stack.Len() != 0 && stack.Back().Value.(Rect).height >= height {
			top := stack.Back().Value.(Rect)
			accumulateWidth += top.width
			ans = max(ans, top.height*accumulateWidth)
			stack.Remove(stack.Back())
		}
		stack.PushBack(Rect{
			accumulateWidth + 1,
			height,
		})
	}
	return ans
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

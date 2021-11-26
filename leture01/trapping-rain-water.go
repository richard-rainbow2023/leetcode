package leture01

import "container/list"

type RainRect struct {
	width  int
	height int
}

func trap(heights []int) int {
	ans := 0
	stack := list.New()
	for _, height := range heights {
		accumulateWidth := 0

		for stack.Len() != 0 && stack.Back().Value.(RainRect).height <= height {
			top := stack.Back().Value.(RainRect)
			bottom := top.height
			accumulateWidth += top.width

			stack.Remove(stack.Back())
			if stack.Len() == 0 {
				continue
			}
			up := min(height, stack.Back().Value.(RainRect).height)
			ans += accumulateWidth * (up - bottom)
		}
		stack.PushBack(RainRect{
			accumulateWidth + 1,
			height,
		})
	}
	return ans
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

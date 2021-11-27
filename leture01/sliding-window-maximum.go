package leture01

import "container/list"

func maxSlidingWindow(nums []int, k int) []int {
	dequeue := list.New()
	var ans []int
	for i := 0; i < len(nums); i++ {
		for dequeue.Len() != 0 && dequeue.Front().Value.(int) <= i-k {
			dequeue.Remove(dequeue.Front())
		}

		for dequeue.Len() != 0 && nums[dequeue.Back().Value.(int)] <= nums[i] {
			dequeue.Remove(dequeue.Back())
		}
		dequeue.PushBack(i)
		if i >= k-1 {
			ans = append(ans, nums[dequeue.Front().Value.(int)])
		}
	}
	return ans
}

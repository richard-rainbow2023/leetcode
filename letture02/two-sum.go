package letture02

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for index, num := range nums {
		if _, ok := dict[target-num]; ok {
			return []int{index, dict[target-num]}
		}
		dict[num] = index
	}
	return nil
}

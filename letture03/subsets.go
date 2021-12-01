package letture03

func subsets(nums []int) (ans [][]int) {
	chosen := []int{}

	var dfs func(int)
	dfs = func(cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), chosen...))
			return
		}
		chosen = append(chosen, nums[cur])
		dfs(cur + 1)
		chosen = chosen[:len(chosen)-1]
		dfs(cur + 1)
	}
	dfs(0)
	return
}

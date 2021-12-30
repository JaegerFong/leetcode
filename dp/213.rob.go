package dp

func rob2(nums []int) int {
	nl := len(nums)
	return max(
		rob2dp(nums, 0, nl-2),
		rob2dp(nums, 1, nl-1))
}

// 环形数组
//
func rob2dp(nums []int, start, end int) int {
	nl := len(nums)
	dp := make([]int, nl+2)
	for i := end; i >= start; i-- {
		dp[i] = max(
			dp[i+1],
			nums[i]+dp[i+2])
	}

	return dp[0]
}

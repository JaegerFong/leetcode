package dp

func rob(nums []int) int {

	nl := len(nums)
	dp := make([]int, nl+2)
	for i := nl - 1; i >= 0; i-- {
		dp[i] = max(
			dp[i+1],
			nums[i]+dp[i+2])
	}

	return dp[nl]
}

// 状态：idx
// 选择: rob || no rob
// dp(nums,idx) = x
// dp(nums,idx) = max( dp(nums,idx+1), nums[idx]+dp(nums,idx+2) )
func robdp(nums []int, start int) int {
	if start > len(nums) {
		return 0
	}

	nr := robdp(nums, start+1)
	r := nums[start] + robdp(nums, start+2)
	res := max(nr, r)
	return res
}

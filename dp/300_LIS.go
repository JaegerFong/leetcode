package dp

// 最长递增子序列
func lengthOfLIS(nums []int) int {

	nlen := len(nums)
	dp := make([]int, nlen)

	for i := 0; i < nlen; i++ {
		dp[i] = 1
	}

	for i := 0; i < nlen; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	max := 1
	for _, v := range dp {
		if v > max {
			max = v
		}
	}

	return max
}

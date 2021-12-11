/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */
package dp

// @lc code=start
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}

	// 求最小值
	res := 0 - (10 * 10 * 10 * 10 * 10 * 10)
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}

	return res
}

// nums
// dp[0]=nums[0]
// dp[1]=max(nums[1],dp[0]+nums[1])
// dp[2]=max(nums[2],dp[1]+nums[2])
// dp[3]=max(nums[3],dp[2]+nums[3])

// dp[i]=max(nums[i],d[i-1]+nums[i])

// dp[0] =-2
// dp[1] = (1,1+-2) = -1
//
// @lc code=end

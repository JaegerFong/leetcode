/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
package dp

func coinChange(coins []int, amount int) int {
	dp := make([]int, 0, amount+1)
	dp[0] = 0
	for i := 0; i < amount+1; i++ {
		// 给予初始值，无穷大。在本例子中，最大值为amt，amt+1则为无穷大
		if i > 0 {
			dp[i] = amount + 1
		}

		for _, c := range coins {
			r := amount - c
			if r < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[r]+1)
		}
	}
}

// @lc code=end

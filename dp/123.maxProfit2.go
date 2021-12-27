package dp

import (
	"fmt"
)

// 限制2次交易,同188
func maxProfit2(prices []int) int {

	lp := len(prices)

	dp := make([][3][2]int, lp)

	for i := 0; i < lp; i++ {
		v := dp[i]
		fmt.Println(v)
	}

	for i := 0; i < lp; i++ {
		for k := 2; k >= 1; k-- {
			latest := i - 1
			if latest == -1 {
				//base case i == 0
				dp[0][k][0] = 0
				dp[0][k][1] = -1000000000

				dp[0][0][0] = 0
				dp[0][0][1] = -1000000000

				continue
			}

			// k-1的含义: 交易次数减少1次，可以在买入时减少或卖出时减少
			dp[i][k][0] = max(dp[latest][k][0], dp[latest][k][1]+prices[i])
			dp[i][k][1] = max(dp[latest][k][1], dp[latest][k-1][0]-prices[i])
		}
	}

	return dp[lp-1][2][0]
}

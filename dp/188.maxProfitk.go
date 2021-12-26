package dp

import "math"

func maxProfitk(k int, prices []int) int {

	pl := len(prices)
	var dp func(int, int) int
	dp = func(start, k int) int {
		if start >= pl {
			return 0
		}

		if k <= 0 {
			return 0
		}

		minBuy := prices[start]
		res := 0
		for i := start + 1; i < pl; i++ {
			minBuy = min(minBuy, prices[i])

			curProfit := prices[i] - minBuy

			restProfit := dp(i+1, k-1)

			res = max(res, curProfit+restProfit)
		}

		return res
	}

	return dp(0, k)
}

// 限制K次交易，最大收益
func maxProfitK(k int, prices []int) int {
	pl := len(prices)
	// 有效买卖次数不应该超过 n
	if k > pl/2 {
		return math.MinInt64
	}
	dp := make([][][2]int, pl) // 初始化三维数组
	for i := 0; i < pl; i++ {
		for j := k; j >= 1; j-- {
			if i-1 == -1 {
				dp[0][k][0] = 0
				dp[0][k][1] = -prices[i]
			}

			dp[i][k][0] = max(dp[i-1][k][0], (dp[i-1][k][1] + prices[i]))
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}

	return dp[pl-1][2][0]
}

// 动态规划+状态机
// 1. 明确状态： N天，K次交易，[0 or 1] 持有或者不持有
// 2. 状态转移： 随着天数减少，持有或不持有时的状态变化
// 3. base case：针对天数和K次交易，分别分析持有和不持有时的base case

// dp定义： dp[N][K][0]=x  dp[N][K][1]=x ，最终取dp[N][0][0]，最后一天股票卖出的利润才是最大的
// dp[N][K][0] = max(dp[N-1][K][0], dp[N-1][K][1] + price[N-1] )
// dp[N][K][1] = max(dp[N-1][K][1], dp[N-1][K-1][0] - price[N-1] )
// 控制变量法，当N=0时
// dp[0][K][0] = 0 , 还没开始，不持有股票，收益为0
// dp[0][K][1] = -infinity ，还没开始，持有股票，不可能的情况
// dp[N][0][0] = 0 , 不允许交易，不持有股票，收益为0
// dp[N][0][1] = -infinity , 不允许交易的情况下，不可能持有股票

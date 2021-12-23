package dp

func maxProfit1(prices []int) int {
	pl := len(prices)
	// TODO: DP备忘录
	var dp func(int) int
	dp = func(start int) int {

		// 数组越界
		if start >= pl {
			return 0
		}

		minBuy := prices[start]
		res := 0
		for i := start + 1; i < pl; i++ {
			minBuy = min(minBuy, prices[start])

			curProfit := prices[i] - minBuy

			// 下一天继续计算
			restProfit := dp(start + 1)

			res = max(res, restProfit+curProfit)
		}

		return res
	}

	return dp(0)
}

// 动态规划：二次循环递归
func mp2dp(prices []int) int {
	pl := len(prices)
	res := 0
	for buy := 0; buy < pl; buy++ {
		for sell := buy + 1; sell < pl; sell++ {
			// 前一天的收益
			beforeMp := mp2dp(prices[sell+1:])
			// 卖出的收益
			profit := prices[sell] - prices[buy]
			res = max(res, beforeMp+profit)
		}
	}

	return res
}

// 最优解，贪心算法
// 核心思想：每天都能比昨天多赚一点
func mp2greedy(prices []int) int {
	pl := len(prices)
	// 最大收益
	profit := 0
	// i从1开始，因为初始时，需要比较i=0的情况，防止越界
	for i := 1; i < pl; i++ {
		// 今天收益
		today := prices[i]
		// 昨天收益
		yesterday := prices[i-1]

		// 今天收益 = 今天价格 - 昨天价格
		todayProfit := today - yesterday

		// 今天收益 > 0，说明有得赚
		if todayProfit > 0 {
			// 今天能赚一点是一点,累积起来
			profit += todayProfit
		}
	}

	return profit
}

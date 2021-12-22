package dp

func maxProfit2() {

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

// 一次循环递归
func mp2dp1(prices []int) int {

	profit := 0
	pl := len(prices)
	minBuy := prices[0]
	for i := 1; i < pl; i++ {
		// 当天买入点和历史最低买入点比较,取最小值
		minBuy := min(minBuy, prices[i])

		// 当天卖价格 - 历史最低买入价格
		curProfit := prices[i] - minBuy

		// 剩余最大收益 = 递归获取i+1后的数据
		restMaxProfit := mp2dp1(prices[i+1:])

		// 取最大收益
		profit = max(profit, curProfit+restMaxProfit)
	}

	return profit
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

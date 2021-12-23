package dp

// 买卖一次，获益最大 => 买入最低，卖出最高
// 每次只需找到最低的买入点，以及计算 (当前卖出点-最低卖出点)和历史最高收益比较
func maxProfit0(prices []int) int {

	// 默认最大收益为0
	profit := 0
	pl := len(prices)
	// 假设第一天为买入的最低价格
	minBuy := prices[0]
	// i=1,从第二天开始遍历
	for i := 1; i < pl; i++ {
		// 当天价格与历史低点比较，取最小值
		minBuy = min(minBuy, prices[i])

		// 当天收益 = 当天价格 - 前段时间的最小买入价格
		curProfit := prices[i] - minBuy

		// 当天收益与历史最大收益，取最大值
		profit = max(profit, curProfit)
	}

	return profit
}

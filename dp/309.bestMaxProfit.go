package dp

func bestMaxProfit(prices []int) int {

	pl := len(prices)

	var dp func(int) int
	dp = func(start int) int {

		if start > pl {
			return 0
		}

		minBuy := prices[start]
		res := 0
		for i := start + 1; i < pl; i++ {
			minBuy = min(minBuy, prices[i])

			curProfit := prices[i] - minBuy
			// 隔一天才能进行买入
			restProfit := dp(i + 2)

			res = max(res, curProfit+restProfit)
		}

		return res
	}

	return dp(0)
}

package dp

func maxProfitWithFee(prices []int, fee int) int {
	pl := len(prices)

	var dp func(int) int
	dp = func(start int) int {
		if start >= pl {
			return 0
		}

		minBuy := prices[start]
		res := 0
		for i := start + 1; i < pl; i++ {
			minBuy = min(minBuy, prices[start])

			curProfit := prices[start] - minBuy - fee

			restProfit := dp(start)

			res = max(res, curProfit+restProfit)
		}

		return res
	}

	return dp(0)
}

package dp

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

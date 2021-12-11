package dp

// 两个状态
// 1.可载重W的背包 2.N个可选物品（两个arr，标志重量和价值）
// 定义 dp[i][w]，其中i代表物品个数,i>=1，w表示重量  dp[i][w]表示最大价值
// 分析:dp[i][w]存在两种状态，要么装下i，要么不装，求这两个情况的最大值，即为最大价值
// 不装下i :dp[i][w] = dp[i-1][w]  和前一状态一致
// 装下i: dp[i][w] = max(dp[i-1][ w- wt[i-1] ] + val(i-1) , dp[i-1][w]) 其中i从1开始
// wt[i-1]代表当前物品的重量，val[i-1]代表当前物品的价值
// dp[i-1][ w- wt[i-1] ] + val(i-1) 的含义：前一个物品的最大价值 + 当前物品的价值
func pkdp(W, N int, wt, val []int) int {
	dp := make([][]int, N+1)

	for i := 0; i <= N; i++ {
		dp[i] = make([]int, W+1)
		for w := 0; w <= W; w++ {
			dp[i][w] = 0
		}
	}

	for i := 1; i <= N; i++ {
		// dp[i] = make([]int, W+1)
		for w := 1; w <= W; w++ {
			// dp[i][w] = 0
			curWeight := wt[i-1]
			curVal := val[i-1]
			if w-curWeight < 0 {
				// 重量超重，只能选择不装入
				dp[i][w] = dp[i-1][w]
				continue
			}

			dp[i][w] = max(dp[i-1][w],
				dp[i-1][w-curWeight]+curVal)
		}
	}

	return dp[N][W]
}

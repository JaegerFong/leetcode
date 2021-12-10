package dp

// 两个状态
// 1.可载重W的背包 2.N个可选物品（两个arr，标志重量和价值）
// 定义 dp[i][w]，其中i代表物品个数，w表示重量  dp[i][w]表示价值
func pkdp(weight, n int, wt, val []int) {
	

	dp := make([][]int,len(n))
	for i := 1; i <= n; i++ {
		witem := make([]int,len(wt))
		dp[i] = witem
		for w := 1;w<=weight;w++ {
			// 容量不足
			if w-wt[i-1] < 0 {
				dp[i][w] = dp[i-1][w]
			}else {
				dp[i][w] = max(dp[i-1][w],
				dp[i-1][w-wt[i-1] + val[i-1])
			}
		}
	}

	return dp[n][weight]
}

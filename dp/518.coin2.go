package dp

// 零钱兑换2
func coinChange11(amount int, coins []int) int {

	cl := len(coins)
	if cl < 0 {
		return 0
	}

	if amount < 0 {
		return 1
	}

	dp := make([][]int, cl+1)
	// init
	for i := 0; i <= cl; i++ {
		item := make([]int, amount+1)
		dp[i] = append(dp[i], item...)
		for j := 0; j <= amount; j++ {
			if j == 0 {
				dp[i][j] = 1
				break
			}
		}
	}

	for i := 1; i <= cl; i++ {
		for j := 1; j <= amount; j++ {
			ret := j - coins[i-1]
			if ret < 0 {
				// 不放入
				dp[i][j] = dp[i-1][j]
				continue
			}

			dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
		}
	}

	return dp[cl][amount]
}

// 等价转换 -> 0-1背包问题
// i个物品,组合成W价值,coint[i]对应i的价值,有多少种组合情况
// 状态: i,amount
// 选择: 放入i或不放入i
// dp[i][j] = x : i物品,j价值,x种方式
// dp[i][j] 可分为放入物品i,或者不放入物品i
// dp[i][j] = d[i-1][j] + dp[i][ j-n[i-1] ]
// base case dp[...][0]=1 价值为0只有1种方式  dp[0][...]=0 物品位0只有0种方式

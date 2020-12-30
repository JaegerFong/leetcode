package dp

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回

// 自己写的
func coinChange(coins []int, amount int) int {
	var dp func([]int, int) int
	dp = func(arr []int, num int) int {
		if num == 0 {
			return 0
		}

		if num < 0 {
			return -1
		}

		var res int
		for _, v := range arr {
			subProblem := dp(arr, num-v)
			if subProblem == -1 {
				continue
			}

			res = min(res, 1+subProblem)
		}

		if res != 0 {
			return res
		} else {
			return -1
		}

	}

	return dp(coins, amount)
}

// 别人的
func coinChange2(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, c := range coins {
			if i < c || dp[i-c] == -1 {
				continue
			}

			count := dp[i-c] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

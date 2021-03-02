package dp

//给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回

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

/*
程序跑不通,栈溢出
暴力递归的思路:
1.base case : amount=0 coin=0 ,amount<0 coin=-1
2.分解子问题:sub_problem = dp(n-coin)
3.选择硬币最少的作为最终结果: min(res,sub_problem+1)
*/
func coinChangeRecurision(coins []int, amount int) int {
	var dp func(int) int
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}

		if amount < 0 {
			return -1
		}

		minCount := 0
		for _, coin := range coins {
			subProblem := dp(amount - coin)
			if subProblem < 0 {
				continue
			}

			minCount = min(minCount, subProblem)
		}

		return minCount
	}

	return dp(amount)
}

/*
stack overflow
递归带备忘录
思路和暴力递归
*/
func coinChangeMemo(coins []int, amount int) int {

	memo := make(map[int]int, amount)
	var dp func(amount int) int
	dp = func(amount int) int {
		if amount == 0 {
			return 0
		}

		if amount < 0 {
			return -1
		}

		var minCount int
		for _, coin := range coins {
			target := amount - coin
			var subProblem int
			if v, ok := memo[target]; ok {
				subProblem = v
			} else {
				subProblem = dp(target)
				if subProblem < 0 {
					continue
				}

				memo[target] = subProblem

			}

			minCount = min(minCount, subProblem)
		}

		return minCount
	}
	return dp(amount)
}

/*
dp数组迭代
*/
func coinChangeIteration(coins []int, amount int) int {
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

package dp

// 416 0-1背包变种，无状态压缩
func cpdp(nums []int) bool {
	var sum int
	for _, v := range nums {
		sum += v
	}

	// 求余不为0,不能平分
	if sum%2 != 0 {
		return false
	}

	sum = sum / 2

	n := len(nums)

	// 初始化dp
	dp := make([][]bool, n+1)

	// 初始化base case
	for i := 0; i <= n; i++ {
		item := make([]bool, sum+1)
		dp[i] = append(dp[i], item...)
		for j := 0; j <= sum; j++ {
			if j == 0 {
				dp[i][j] = true
				break
			}
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j-nums[i-1]]
			}

		}
	}

	return dp[n][sum]
}

// state: N个物品 重量W
// choose: 是否装入背包,true or false
// dp[i][j]=x : i表示i个物品,j表示j重量,x表示是否装入背包
// base case: dp[...][0]=true(没有空间,相当于装满)   dp[0][...]=false(没有物品，永远也装不满)
// case 1: 如果不放入物品,则状态和前一个状态一致,dp[i][j]=dp[i-1][j]
// case 2: 如果放入物品,则 dp[i][j] = dp[i-1][ j-num[i-1] ]

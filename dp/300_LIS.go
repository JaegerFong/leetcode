package dp

// 最长递增子序列
func lengthOfLIS(nums []int) int {

	nlen := len(nums)
	dp := make([]int, nlen)

	for i := 0; i < nlen; i++ {
		dp[i] = 1
	}

	for i := 0; i < nlen; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	max := 1
	for _, v := range dp {
		if v > max {
			max = v
		}
	}

	return max
}

func lis(arr []int) int {
	l := len(arr)
	dp := make([]int, l)
	dp[0] = 1
	for i := 1; i < l; i++ {
		if arr[i] > arr[i-1] {
			dp[i] = dp[i-1] + 1
			continue
		}

		dp[i] = dp[i-1]
	}

	return dp[l-1]
}

// arr = [1,0,2,4,3,6]
// i = len(arr)-1
// dp[i] => arr[i-1] > arr[i] ? dp[i-1] : dp[i-1]+1

// dp[5] => arr[4] > arr[5] ? dp[4] : dp[4]+1 => 4
// dp[4] => arr[3] > arr[4] ? dp[3] : dp[3]+1 => 3
// dp[3] => arr[2] > arr[3] ? dp[2] : dp[2]+1 => 3
// dp[2] => arr[1] > arr[2] ? dp[1] : dp[1]+1 => 2
// dp[1] => arr[0] > arr[1] ? dp[0] : dp[0]+1 => 1
// dp[0] => 1

package dp

func fib2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fib2(n-1) + fib2(n-2)
}

// fib(1) = 1
// fib(2) = 1
// fib(3) = fib(1) + fib(2) = 2
var tmpM = make(map[int]int)

func fib2m(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if r, ok := tmpM[n]; ok {
		return r
	}

	tmpM[n] = fib2m(n-1) + fib2m(n-2)

	return tmpM[n]
}

func fib2dp(n int) int {
	tm := make(map[int]int, n)

	tm[1], tm[2] = 1, 1
	for i := 3; i <= n; i++ {
		tm[i] = tm[i-1] + tm[i-2]
	}

	return tm[n]
}

var coins []int

// 递归
// @param amt
// @return int count
func dp(amt int) int {
	if amt == 0 {
		return 0
	}

	if amt < 0 {
		return -1
	}

	var basemin int
	for idx, v := range coins {
		// choose v
		if idx == 0 {
			basemin = dp(amt-v) + 1
		}

		basemin = min(basemin, dp(amt-v)+1)
	}

	return 0
}

// 迭代
func coinchange(coins []int, amt int) int {

	if amt == 0 {
		return 0
	}

	// 无解
	if amt < 0 {
		return -1
	}

	dp := make([]int, amt+1)
	dp[0] = 0
	for i := 0; i < amt+1; i++ {
		// 给予初始值，无穷大。在本例子中，最大值为amt，amt+1则为无穷大
		if i > 0 {
			dp[i] = amt + 1
		}

		for _, c := range coins {
			r := i - c
			if r < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[r]+1)
		}
	}

	// 没有得出最小值，取到的是无穷大的值
	if dp[amt] == amt+1 {
		return -1
	}

	return dp[amt]
}

// coin : 1,2,5
// dp(10) => min{dp(9),dp(8),dp(5)} + 1
// dp(9) => min{dp(8),dp(7),dp(5)} + 1
// dp(8) => min{dp(7),dp(6),dp(4)} + 1

// dp(5) => min(dp(4)+1 ,dp(3)+1 ,dp(0)+1) => min(3,3,1) = 1  取除了0之外的最小值
// dp(4) => min(dp(3)+1 ,dp(2)+1 ,dp(-1)+1) => min(3,2,0) = 2 取除了0之外的最小值
// dp(3) => min(dp(2)+1 ,dp(1)+1 ,dp(-2)+1) => min(2,2,0) = 2 取除了0之外的最小值
// dp(2) => min(dp(1)+1 ,dp(0)+1 ,dp(-3)+1) => min(2,1,0) = 1  取除了0之外的最小值
// dp(1) => min(dp(0)+1 ,dp(-1)+1,dp(-4)+1) => min(1,0,0) = 1  取除了0之外的最小值
// base case : dp(0) = 0 ,dp(<0) = -1  => dp(1) dp(2) ... dp(10)

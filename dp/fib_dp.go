package dp

// 递归解斐波那契数列,存在冗余子问题,子问题个数为2^n,即时间复杂度O(2^n)
func fib_recurision(n int) int {
	if n == 1 || n == 2 {
		return 0
	}
	return fib_recurision(n-1) + fib_recurision(n-2)
}

// 带备忘录解斐波那契数列 , 本算法不存在冗余子问题,子问题个数为n,即时间复杂度O(n)
func fib_memo(n int) int {
	if n < 1 {
		return 0
	}

	memo := make(map[int]int, n)

	return fibm(memo, n)
}
func fibm(m map[int]int, n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	if v, ok := m[n]; ok {
		return v
	}

	return fibm(m, n-1) + fibm(m, n-2)
}

// dp table解斐波那契数列,迭代
func fib_dp_table(n int) int {
	if n == 1 || n == 2 {
		return 0
	}

	dp_table := make([]int, n+1)
	dp_table[1], dp_table[2] = 1, 1
	for i := 3; i <= n; i++ {
		dp_table[i] = dp_table[i-1] + dp_table[i-2]
	}

	return dp_table[n]
}

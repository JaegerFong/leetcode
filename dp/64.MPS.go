package dp

func minPathSum(grid [][]int) int {

	i := 0 // 一维数组长度
	j := 0 // 二维数组长度

	return mpsdp(grid, i, j)
}

func mpsdp(grid [][]int, i, j int) int {

	// base case

	return min(mpsdp(grid, i-1, j), mpsdp(grid, i, j-1)) + grid[i][j]
}

// state: location
// choose: 向左或向下,从i,j => min ([i+1][j] || [i][j+1])
// dp(i,j) => min(dp[i-1][j],dp[i][j-1]) + grid[i][j]

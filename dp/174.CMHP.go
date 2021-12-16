package dp

// caculateMinHP()
func CMHP(grid [][]int) int {

	return 0
}

func cmhpdp(grid [][]int, i, j int) int {
	m := len(grid)    // x
	n := len(grid[0]) // y

	// base case
	if i == m-1 && j == n-1 {
		if grid[i][j] > 0 {
			return 1
		}

		return -grid[i][j] + 1
	}

	res := min(cmhpdp(grid, i, j+1), cmhpdp(grid, i+1, j)) - grid[i][j]

	if res <= 0 {
		return 1
	}

	return res
}

// dp定义: dp(i,j) = x, 其中i,j代表路径，x表示最小血量
// 已知走到 dp(i,j+1)时的生命值,减去上一步的生命值 grid(i)(j) 可知最小的生命值
// base case 1: 如果减去grid[i][j],为负数,则代表grid[i][j] > grid[i][j+1],则应该取1
// base case 2: 如果为正数,则代表grid[i][j]<grid[i][j+1],返回差值,即dp(i,j)的最小生命值为res
// res = min(dp[i+1][j],dp[i][j+1])-grid[i][j]
// dp(i,j) = res<=0 ? 1 : res

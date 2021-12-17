package dp

// caculateMinHP()
var memo [][]int

func CMHP(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	memo = make([][]int, m)

	for i := 0; i < m; i++ {
		item := make([]int, n)
		memo[0] = append(memo[0], item...)
		for j := 0; j < n; j++ {
			memo[i][j] = 0
		}
	}

	return cmhpdp2(grid, 0, 0)
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

func cmhpdp2(grid [][]int, i, j int) int {
	m := len(grid)
	n := len(grid[0])
	// base case
	if m-1 == i && n-1 == j {
		if grid[i][j] > 0 {
			memo[i][j] = 1
			return memo[i][j]
		}
		memo[i][j] = -grid[i][j] + 1
		return memo[i][j]
	}

	res := min(cmhpdp(grid, i+1, j), cmhpdp(grid, i, j+1)) - grid[i][j]

	if res <= 0 {
		memo[i][j] = 1
		return memo[i][j]
	}

	memo[i][j] = 0
	return memo[i][j]
}

//

// dp函数定义：从grid[i][j] 走到终点所需的最小血量
// dp定义: dp(i,j) = x, 其中i,j代表路径，x表示最小血量
// 其中 x = (min[dp(i+1,j),dp(i,j+1)] - grid[i][j] ) <=0 ? 1: (min[dp(i+1,j),dp(i,j+1)] - grid[i][j] ) <=0
// 如何解释?
// dp(i,j) 代表最小血量，走到dp(i+1,j)时的血量减去上一步的血量,如果为负数，则代表grid[i][j] > dp(i+1,j),
// 已知走到 dp(i,j+1)时的生命值,减去上一步的生命值 grid(i)(j) 可知最小的生命值
// base case 1: 如果减去grid[i][j],为负数,则代表grid[i][j] > grid[i][j+1],则应该取1
// base case 2: 如果为正数,则代表grid[i][j]<grid[i][j+1],返回差值,即dp(i,j)的最小生命值为res
// res = min(dp[i+1][j],dp[i][j+1])-grid[i][j]
// dp(i,j) = res<=0 ? 1 : res

// dp(i,j) => min[dp(i+1,j),dp(i,j+1)]
// dp(0,0) => min[dp(1,0),dp(0,1)]
// 	dp(1,0) => min[dp(2,0),dp(1,1)]
// 	dp(0,1) => min[dp(1,1),dp(0,2)]
//	  dp(2,0) => min(dp(3,0),dp(2,1))
//	  dp(1,1) => min(dp(2,1),dp(1,2))
// 	  dp(0,2) => min(dp(1,2),dp(0,3))

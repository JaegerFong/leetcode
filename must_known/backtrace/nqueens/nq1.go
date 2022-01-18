package nqueens

import "strings"

func solveNQueens(n int) [][]string {
	bd := make([][]string, n)
	for i := range bd {
		bd[i] = make([]string, n)
		for j := range bd[i] {
			bd[i][j] = "."
		}
	}
	res := [][]string{}
	helper(0, bd, &res, n)
	return res
}

// 选择
// 约束
// 回溯目标
func helper(r int, bd [][]string, res *[][]string, n int) {
	// 到达最后一行
	if r == n {
		temp := make([]string, len(bd))
		for i := 0; i < n; i++ {
			temp[i] = strings.Join(bd[i], "")
		}
		*res = append(*res, temp)
		return
	}

	//
	for c := 0; c < n; c++ {
		// 只进入满足条件的选项
		if isValid(r, c, n, bd) {
			// 选择
			bd[r][c] = "Q"
			// 追溯
			helper(r+1, bd, res, n)
			// 回撤
			bd[r][c] = "."
		}
	}

}

// 排除：同一列，左上角，右上角存在Q，则返回 false
func isValid(r, c int, n int, bd [][]string) bool {
	for i := 0; i < r; i++ {
		for j := 0; j < n; j++ {
			if bd[i][j] == "Q" && (j == c || i+j == r+c || i-j == r-c) {
				return false
			}
		}
	}
	return true
}

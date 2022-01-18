package nqueens

func solveNQueens(n int) [][]string {

}

// 路径：记录track
// 选择列表：剩余的可选项
// 结束条件：没有可选项了
func backtrack(board [][]string, row int) {
	lr := len(board)

	// 结束
	if row == lr {
		return
	}

	lc := len(board[row])

	for col := 0; col < lc; col++ {
		// 排除不合法选择
		if !checkNQValid(board, row, col) {
			continue
		}

		// 选择
		board[row][col] = "Q"

		// 进入下一决策
		backtrack(board, row+1)

		// 撤销选择
		board[row][col] = "."
	}
}

func checkNQValid(board [][]string, row, col int) bool {

	// 检查当前列是否有Q

	// 检查右上方

	// 检查左上方

	return true
}

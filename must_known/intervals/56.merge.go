package intervals

func merge(intervals [][]int) [][]int {
	// 按起点升序排序

	left := intervals[0][0]
	right := intervals[0][1]
	res := make([][]int, 1)
	for _, v := range intervals {
		curLeft := v[0]
		curRight := v[1]

		// 区间内，跳过

		// 有交集，合并
		if curLeft <= right && curRight >= right {
			right = curRight
		}

		// 区间外，保存当前所合并的区间，并更新区间范围
		if curLeft > right {
			curRange := []int{left, right}
			res = append(res, curRange)

			left = curLeft
			right = curRight
		}

	}

	return res
}

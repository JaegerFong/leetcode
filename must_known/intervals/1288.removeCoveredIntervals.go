package intervals

func removeCoveredIntervals(intervals [][]int) int {

	// 起点升序排序

	// 初始起点
	// 初始终点
	left := intervals[0][0]
	right := intervals[0][1]
	res := 0
	for _, v := range intervals {

		curLeft := v[0]
		curRight := v[1]
		// 位于区间中，被覆盖
		if curLeft >= left && curRight <= right {
			res++
		}

		// 区间相交，合并区间
		if curLeft <= right && curRight >= right {
			right = curRight
		}

		// 区间之外，重新定义区间
		if curLeft > right {
			left = curLeft
			right = curRight
		}

	}

	return res
}

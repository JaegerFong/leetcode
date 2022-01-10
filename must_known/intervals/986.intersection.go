package intervals

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	// 双指针
	i, j := 0, 0

	res := make([][]int, 1)
	for i < len(firstList) && j < len(secondList) {
		a1 := firstList[i][0]
		a2 := firstList[i][1]

		b1 := secondList[j][0]
		b2 := secondList[j][1]

		// 存在区间
		// a1 > b2 || a2 < b1
		// b2 >= a1 && a2 >= b1
		if b2 >= a1 && a2 >= b1 {
			// 取最值
			start := max(a1, b1)
			end := min(a2, b2)
			item := []int{start, end}
			res = append(res, item)
		}

		// 判断移动哪一个指针
		if b2 < a2 {
			i++
			continue
		}

		j++
	}

	return res
}

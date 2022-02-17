package hash

func isPossible(nums []int) bool {
	freq := make(map[int]int)
	need := make(map[int]int)

	for idx, _ := range nums {
		freq[idx]++
	}

	for idx, _ := range nums {
		// 已被其他序列使用
		if freq[idx] == 0 {
			continue
		}

		// 先判断是否能接到其他子序列后面
		if need[idx] > 0 {
			freq[idx]--
			need[idx]--
			need[idx+1]++
		} else if freq[idx] > 0 && freq[idx+1] > 0 && freq[idx+2] > 0 {
			freq[idx]--
			freq[idx+1]--
			freq[idx+2]--
			// 对后值有需求
			need[idx+3]++
		} else {
			return false
		}

	}

	return true
}

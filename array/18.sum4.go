package array

func fourSum(nums []int, target int) [][]int {
	return nSum(nums, 4, 0, target)
}

func nSum(nums []int, n, start, target int) [][]int {

	sz := len(nums)
	if n < 2 || sz < n {
		return nil
	}

	res := make([][]int, 1)
	if n == 2 {
		// 求2sum
		left := nums[start]
		right := nums[sz-1]
		sum := left + right
		for start < sz {
			if sum < target {
				for start < sz && left == nums[start] {
					start++
				}
			} else if sum > target {
				for sz > start && right == nums[sz] {
					sz--
				}
			} else {
				res = append(res, []int{start, sz})
				for start < sz && left == nums[start] {
					start++
				}

				for sz > start && right == nums[sz] {
					sz--
				}
			}
		}

	} else {
		// 求 nsum
		for i := start; i < sz; i++ {
			sumn := nSum(nums, n-1, i+1, target)
			for _, v := range sumn {
				v = append(v, nums[i])
				res = append(res, v)
			}

			for i < sz-1 && nums[i] == nums[i+1] {
				i++
			}
		}
	}

	return res

}

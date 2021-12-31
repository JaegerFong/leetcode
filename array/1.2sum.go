package array

// 双指针
func twoSumn(nums []int, target int) []int {
	nl := len(nums)
	l := 0
	r := nl - 1
	for l < r {
		s := nums[l] + nums[r]
		if s < target {
			l++
			continue
		} else if s > target {
			r--
			continue
		} else {
			break
		}
	}

	return []int{l, r}
}

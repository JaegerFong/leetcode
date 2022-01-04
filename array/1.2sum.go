package array

// 双指针
func twoSumn(nums []int, target int) []int {
	nl := len(nums)
	l := 0
	r := nl - 1

	for l < r {
		left := nums[l]
		right := nums[r]
		sum := left + right

		if sum < target {

			// 跳过相同的数
			for l < r && nums[l] == left {
				l++
			}

		} else if sum > target {

			// 跳过相同的数
			for l < r && nums[r] == right {
				r--
			}

		} else {

			// 跳过相同的数
			for l < r && nums[l] == left {
				l++
			}

			// 跳过相同的数
			for l < r && nums[r] == right {
				r--
			}

			break
		}
	}

	return []int{l, r}
}

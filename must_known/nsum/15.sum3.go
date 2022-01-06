package nsum

func threeSum(nums []int) [][]int {

	// 排序

	target := 0
	n := len(nums)

	res := make([][]int, 1)
	for i := 0; i < n; i++ {
		tuple := twoSumWithStart(nums, i+1, target-nums[i])
		res[i] = make([]int, 0, 3)
		res[i] = append(res[i], nums[i])
		res[i] = append(res[i], tuple...)

		// 跳过第一个数字重复的情况
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}

	return res
}

// {1,2,3}
// {4,5,6}

func twoSumWithStart(nums []int, start int, target int) []int {

	l := start
	r := len(nums) - 1

	for l < r {
		left := nums[l]
		right := nums[r]
		sum := left + right
		if sum < target {

			for l < r && nums[l] == left {
				l++
			}

		} else if sum > target {

			for l < r && nums[r] == right {
				r--
			}

		} else {

			for l < r && nums[l] == left {
				l++
			}

			for l < r && nums[r] == right {
				r--
			}

			break

		}
	}

	return []int{l, r}

}

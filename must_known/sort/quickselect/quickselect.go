package quickselect

func partition(nums []int, left, right int) int {
	if left == right {
		return left
	}

	pivot := nums[left]

	i := left
	j := right + 1

	for true {

		// 保证 nums[left...i] 小于 pivot
		for true {

			// 大于pivot，则退出循环
			i++
			curL := nums[i]
			if curL > pivot {
				break
			}

			// 到最后的节点了
			if i == right {
				break
			}

		}

		// 保证nums[j...right] 都大于pivot
		for true {

			j--
			curR := nums[j]
			if curR < pivot {
				break
			}

			if j == left {
				break
			}

		}

		if i >= j {
			break
		}

		// 存在 nums[i] > pivot && nums[j] < pivot
		// 需要交换位置
		swap(nums, i, j)
	}

	swap(nums, j, left)

	return 0
}

func swap(nums []int, i, j int) []int {
	nums[j], nums[i] = nums[i], nums[j]
	return nums
}

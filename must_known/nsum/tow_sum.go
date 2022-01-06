package nsum

// 两数之和
// 解法一：暴力循环
// 时间复杂度：O（n^2）
func twoSum(nums []int, target int) []int {
	count := len(nums)
	var idxs []int
	for idx, val := range nums {
		for i := idx + 1; i < count; i++ {
			if val+nums[i] == target {
				idxs = []int{
					idx, i,
				}

				return idxs
			}
		}
	}
	return nil
}

// 解法二：哈希表辅助
func twoSum2(nums []int, target int) []int {
	targetMap := map[int]int{}
	for idx, val := range nums {
		if i, ok := targetMap[target-val]; ok {
			return []int{
				i, idx,
			}
		}
		targetMap[val] = idx
	}
	return nil
}

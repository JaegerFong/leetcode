package divide

import "strconv"

func diffWaysToCompute(expression string) []int {
	result := make([]int, 0)
	for idx, _ := range expression {
		c := expression[idx]
		if c == '+' || c == '*' || c == '-' {
			// 分
			left := diffWaysToCompute(expression[0:idx])
			right := diffWaysToCompute(expression[idx+1:])
			// 治
			for _, l := range left {
				for _, r := range right {
					if c == '+' {
						result = append(result, l+r)
						continue
					}

					if c == '-' {
						result = append(result, l-r)
						continue
					}

					if c == '*' {
						result = append(result, l*r)
						continue
					}
				}
			}
		}
	}

	// base case
	if len(result) == 0 {
		i, _ := strconv.Atoi(expression)
		result = append(result, i)
	}

	return result
}

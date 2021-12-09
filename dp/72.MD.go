package dp

func MD(s1, s2 string) int {

	return 0
}

// 涉及到两字符串比对，优先考虑双指针i,j
// dp(i,j)
// dpm[i][j]
// dpm[i-1][j-1] dpm[i][j-1] dpm[i-1][j]
// i=3 j=4
// dpm[0][0] dpm[0][1] dpm[0][2] dpm[0][3]
// dpm[]
// recursion
func mddp(s1, s2 string, i, j int) int {

	// base case i指针走完，则返回 j剩余的长度
	if i == -1 {
		return j + 1
	}

	// j指针走完，则返回 i剩余的长度
	if j == -1 {
		return i + 1
	}

	if s1[i] == s2[j] {
		return mddp(s1, s2, i-1, j-1)
	}

	return min3(
		mddp(s1, s2, i, j-1),   // 插入
		mddp(s1, s2, i-1, j),   // 删除
		mddp(s1, s2, i-1, j-1), // 替换
	)
}

func min3(i, j, k int) int {
	if i > j && i > k {
		return i
	}

	if j > i && j > k {
		return j
	}

	return k
}

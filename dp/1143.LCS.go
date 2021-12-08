package dp

var lcsmap [][]int

func LCS(text1, text2 string) int {
	i, j := len(text1), len(text2)
	lcsmap = make([][]int, i)
	for n := 0; n < i; n++ {
		lcsmap[n] = make([]int, j)
	}

	return lcsdp(text1, text2, 0, 0)
}

// base case i == len(s1) && j == len(s2) return 0
// s1[i] == s[j] => 进行下一位的比较 s[i+1] s[j+1]
// s1[i] != s[j] => 考虑多种情况
// 1.s1[i] 不在公共子序列中 i+1,j
// 2.s2[j] 不在公共子序列中 i,j+1
// 3.s1[i] s2[j] 都不在公共子序列中,移到下一位 i+1,j+1 （忽略情况3,本题求最长，那剩余的字串不可能比目前的位置公共子序列还大）
func lcsdp(s1, s2 string, i, j int) int {

	if i == len(s1) && j == len(s2) {
		return 0
	}

	if lcsmap[i][j] != 0 {
		return lcsmap[i][j]
	}

	if s1[i] == s2[j] {
		lcsmap[i][j] = 1 + lcsdp(s1, s2, i+1, j+1)
	} else {
		lcsmap[i][j] = max(lcsdp(s1, s2, i+1, j), lcsdp(s1, s2, i, j+1))
	}

	return lcsmap[i][j]
}

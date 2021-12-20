package dp

func im(s string, p string) bool {

	return false
}

func imdp(s, p string, i, j int) bool {

	// p[j]: .
	// p[j+1]: *
	// s[i],p[j+2]: 代表跳过p[j]字符,p[j+1]*,匹配0次
	// s[i+1]: 下一字符
	// p[j]当前*匹配
	if s[i] == p[j] || p[j] == '.' {
		// 匹配
		if j < len(p)-1 && p[j+1] == '*' {
			// 存在*,匹配0或多次
			return imdp(s, p, i, j+2) || imdp(s, p, i+1, j)
		} else {
			return imdp(s, p, i+1, j+1)
		}

	} else {
		// 不匹配
		if j < len(p)-1 && p[j+1] == '*' {
			// 存在*,匹配0
			return imdp(s, p, i, j+2)
		} else {
			return false
		}
	}

}

// 两字符串比对,必先考虑使用双指针i,j
// dp定义: dp(s,p string i,j int) = x
// i,j: 再s,p上的是位置
// x: 是否匹配
// s[i]  p[j] 匹配
// s[i] == p[j] || p[j]== '.' : 字符相等或匹配模式为单个字符匹配
//	1.1 j<len(p)-1 && p[j+1] == '*' j+1的字符为*
// 		case 1: 匹配0次 i,j+2
// 		case 2: 匹配多次 i+1,j
//	1.2 下一字符不为*
//		i++;j++;
//	2.1 j < len(p)-1 && p[j+1] == '*' j+1的字符为*
//		case 1: j+2
//	2.2 s[i] != p[j] && p[j] != '.' && p[j+1] != '*'
//		return false

func imdp2(s, p string, i, j int) bool {
	n := len(p)
	if s[i] == p[j] || p[j] == '.' {
		if j < n-1 && p[j+1] == '*' {
			return imdp2(s, p, i, j+2) ||
				imdp2(s, p, i+1, j)
		} else {
			return imdp2(s, p, i+1, j+1)
		}
	} else {
		if j < n-1 && p[j+1] == '*' {
			return imdp2(s, p, i, j+2)
		} else {
			return false
		}
	}

}

//
